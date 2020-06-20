package auth

import (
	"errors"
	"go-boilerplate/pkg/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Error constants
var (
	ErrAuthFailed = errors.New("Client Authentication failed")
)

// IRepository IRepository
type IRepository interface {
	FindUser(id uint) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
}

// Service Service
type Service struct {
	rep IRepository
}

// NewService New Service
func NewService(repository IRepository) *Service {
	return &Service{
		rep: repository,
	}
}

// Attempt 驗證帳號密碼
func (s *Service) Attempt(username string, password string) (u *models.User, err error) {
	u, err = s.rep.FindUserByUsername(username)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err == nil {
		return
	}

	return nil, ErrAuthFailed
}

// GetUserFromToken GetUserFromToken
func (s *Service) GetUserFromToken(tokenString string) (user *models.User, err error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return
	}

	if token.Valid {
		id, _ := strconv.Atoi(token.Claims.(*jwt.StandardClaims).Subject)
		user, err = s.rep.FindUser(uint(id))
		if err != nil {
			return
		}
	}

	return
}
