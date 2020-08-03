package services

import (
	"errors"
	"fmt"
	"go-boilerplate/internal/app/models"
	"go-boilerplate/pkg/auth"
	"go-boilerplate/pkg/config"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Error constants
var (
	ErrAuthFailed = errors.New("Client Authentication failed")
)

// IUserRepository IUserRepository
type IUserRepository interface {
	Create(user models.User) (*models.User, error)
	Update(id uint, user models.User) (*models.User, error)
	Find(id uint) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
}

// AuthService AuthService
type AuthService struct {
	rep IUserRepository
}

// NewAuthService New AuthService
func NewAuthService(repository IUserRepository) *AuthService {
	return &AuthService{
		rep: repository,
	}
}

// Attempt 驗證帳號密碼
func (s *AuthService) Attempt(username string, password string) (user *models.User, err error) {
	user, err = s.rep.FindByUsername(username)
	if err != nil {
		return
	}

	if user.EmailVerifiedAt == nil {
		return nil, ErrAuthFailed
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		return
	}

	return nil, ErrAuthFailed
}

// GetUserFromToken GetUserFromToken
func (s *AuthService) GetUserFromToken(tokenString string) (user *models.User, err error) {
	u, err := s.GetUserFromUnverified(tokenString)
	if err != nil {
		return
	}

	token, err := auth.ParseToken(tokenString, u.Password)
	if err != nil {
		return
	}

	if token.Valid {
		user = u
	}

	return
}

// GetUserFromUnverified GetUserFromUnverified
func (s *AuthService) GetUserFromUnverified(tokenString string) (user *models.User, err error) {
	claims, err := auth.Decode(tokenString)
	if err != nil {
		return nil, err
	}

	id, _ := strconv.Atoi(claims.Subject)
	user, err = s.rep.Find(uint(id))

	return
}

// Register Register
func (s *AuthService) Register(user models.User) (*models.User, error) {
	return s.rep.Create(user)
}

// EmailVerify EmailVerify
func (s *AuthService) EmailVerify(id uint) (*models.User, error) {
	t := time.Now()
	return s.rep.Update(id, models.User{
		EmailVerifiedAt: &t,
	})
}

// SendEmailVerification SendEmailVerification
func (s *AuthService) SendEmailVerification(username string) (string, error) {
	user, err := s.rep.FindByUsername(username)
	if err != nil {
		return "", err
	}

	token, _, err := auth.CreateToken(user.ID, user.Password)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/confirm-email?id=%d&signature=%s", config.App.Server.Host, user.ID, token), nil
}

// SendResetLink SendResetLink
func (s *AuthService) SendResetLink(username string) (string, error) {
	user, err := s.rep.FindByUsername(username)
	if err != nil {
		return "", err
	}

	token, _, err := auth.CreateToken(user.ID, user.Password)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/reset-password?token=%s", config.App.Server.Host, token), nil
}

// PasswordReset PasswordReset
func (s *AuthService) PasswordReset(tokenString string, password string) (*jwt.StandardClaims, error) {
	claims, err := auth.Decode(tokenString)
	if err != nil {
		return nil, err
	}

	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return nil, err
	}

	user, err := s.rep.Find(uint(userID))
	if err != nil {
		return nil, err
	}

	_, err = auth.ParseToken(tokenString, user.Password)
	if err != nil {
		return nil, err
	}

	s.rep.Update(uint(userID), models.User{
		Password: password,
	})

	return claims, nil
}
