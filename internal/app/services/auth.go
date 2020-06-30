package services

import (
	"errors"
	"go-boilerplate/internal/app/models"
	"go-boilerplate/pkg/auth"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Error constants
var (
	ErrAuthFailed = errors.New("Client Authentication failed")
)

// IUserRepository IUserRepository
type IUserRepository interface {
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

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		return
	}

	return nil, ErrAuthFailed
}

// GetUserFromToken GetUserFromToken
func (s *AuthService) GetUserFromToken(tokenString string) (user *models.User, err error) {
	token, err := auth.ParseToken(tokenString)
	if err != nil {
		return
	}

	if token.Valid {
		id, _ := strconv.Atoi(token.Claims.(*jwt.StandardClaims).Subject)
		user, err = s.rep.Find(uint(id))
		if err != nil {
			return
		}
	}

	return
}
