package auth

import (
	"errors"
	"go-boilerplate/pkg/config"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Error constants
var (
	ErrTokenExpired = errors.New("Token has expired and can no longer be refreshed")
)

// CreateToken Create JWT
func CreateToken(id uint) (token string, expire time.Time, err error) {
	expire = time.Now().Add(config.App.TTL * time.Second)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Id:        uuid.New().String(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    config.Server.Host,
		NotBefore: time.Now().Unix(),
		Subject:   strconv.Itoa(int(id)),
		ExpiresAt: expire.Unix(),
	})

	token, err = claims.SignedString([]byte(config.App.Key))

	return
}

// RefreshToken Refresh JWT
func RefreshToken(tokenString string) (refreshToken string, err error) {
	token, err := ParseToken(tokenString)

	var validationError *jwt.ValidationError
	if errors.As(err, &validationError) && validationError.Errors == jwt.ValidationErrorExpired {
		if payload, ok := token.Claims.(*jwt.StandardClaims); ok {
			refreshTTL := time.Unix(payload.IssuedAt, 0).Add(config.App.RefreshTTL * time.Second).Unix()
			if time.Now().Unix() > refreshTTL || payload.IssuedAt > time.Now().Unix() {
				return "", ErrTokenExpired
			}

			expiresAt := time.Now().Add(config.App.TTL * time.Second).Unix()
			if expiresAt > refreshTTL {
				expiresAt = refreshTTL
			}

			claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
				Id:        uuid.New().String(),
				IssuedAt:  payload.IssuedAt,
				Issuer:    payload.Issuer,
				NotBefore: time.Now().Unix(),
				Subject:   payload.Subject,
				ExpiresAt: expiresAt,
			})

			refreshToken, err = claims.SignedString([]byte(config.App.Key))
		}
	}

	return
}

// ParseToken Parse JWT
func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.App.Key), nil
	})
}
