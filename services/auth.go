package services

import (
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fantasy9830/go-boilerplate/repositories"
	"golang.org/x/crypto/bcrypt"
)

// AuthService ...
type AuthService struct {
	SigningKey     []byte
	userRepository *repositories.UserRepository
}

// Claims ...
type Claims struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	jwt.StandardClaims
}

// NewAuthService constructor
func NewAuthService() *AuthService {
	return &AuthService{
		SigningKey:     []byte(conf.GetString("signingKey")),
		userRepository: repositories.NewUserRepository(),
	}
}

// Attempt 驗證帳號密碼
func (service *AuthService) Attempt(username string, password string) bool {
	user := service.userRepository.Lookup(username)

	err := bcrypt.CompareHashAndPassword([]byte(user.Secret), []byte(password))

	return err == nil
}

// GenerateToken 產生token
func (service *AuthService) GenerateToken(username string) (token string, err error) {
	user := service.userRepository.Lookup(username)

	claims := Claims{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Address:  user.Address,
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.Itoa(int(user.ID)),
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			Issuer:    "Ricky",
			IssuedAt:  time.Now().Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = tokenClaims.SignedString(service.SigningKey)

	return
}

// ParseToken 解析token
func (service *AuthService) ParseToken(token string) (*Claims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return service.SigningKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// GetRoles 取得roles
func (service *AuthService) GetRoles(userID uint, guardName string) []string {
	roles := service.userRepository.Roles(userID, guardName)

	var result []string
	for _, value := range roles {
		result = append(result, value.Name)
	}

	return result
}

// GetPermissions 取得permissions
func (service *AuthService) GetPermissions(userID uint, guardName string) map[string][]string {
	permissions := service.userRepository.Permissions(userID, guardName)

	result := make(map[string][]string)
	for _, value := range permissions {
		result[value.Action] = append(result[value.Action], value.Name)
	}

	return result
}
