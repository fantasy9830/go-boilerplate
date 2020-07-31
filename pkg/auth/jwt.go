package auth

import (
	"encoding/json"
	"errors"
	"go-boilerplate/pkg/config"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Error constants
var (
	ErrTokenExpired = errors.New("Token has expired and can no longer be refreshed")
	ErrTokenInvalid = errors.New("Token is invalid")
)

// CreateToken Create JWT
func CreateToken(id uint, key string) (token string, expire time.Time, err error) {
	expire = time.Now().Add(config.App.TTL * time.Second)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Id:        uuid.New().String(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    config.Server.Host,
		NotBefore: time.Now().Unix(),
		Subject:   strconv.Itoa(int(id)),
		ExpiresAt: expire.Unix(),
	})

	token, err = claims.SignedString([]byte(key))

	return
}

// RefreshToken Refresh JWT
func RefreshToken(tokenString string, key string) (refreshToken string, err error) {
	token, err := ParseToken(tokenString, key)

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

			refreshToken, err = claims.SignedString([]byte(key))
		}
	}

	return
}

// ParseToken Parse JWT
func ParseToken(tokenString string, key string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
}

// Decode Decode
func Decode(tokenString string) (claims *jwt.StandardClaims, err error) {
	payload := strings.Split(tokenString, ".")
	if len(payload) != 3 {
		return nil, ErrTokenInvalid
	}

	bytes, err := jwt.DecodeSegment(payload[1])
	if err != nil {
		return nil, err
	}

	json.Unmarshal(bytes, &claims)

	return
}
