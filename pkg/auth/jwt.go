package auth

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrTokenInvalid        = errors.New("token is invalid")
	ErrRefreshTokenExpired = errors.New("token has expired and can no longer be refreshed")
)

type Claims struct {
	jwt.RegisteredClaims
}

type JWTer interface {
	CreateClaims(userID string, issuer string) Claims
	CreateToken(claims Claims) (string, error)
	ParseToken(token string) (*jwt.Token, error)
	RefreshToken(token string) (string, error)
	Decode(token string) (*Claims, error)
	TTL() time.Duration
	RefreshTTL() time.Duration
}

type JWT struct {
	key        []byte
	ttl        time.Duration
	refreshTTL time.Duration
}

func WithTTL(duration time.Duration) func(*JWT) {
	return func(j *JWT) {
		j.ttl = duration
	}
}

func WithRefreshTTL(duration time.Duration) func(*JWT) {
	return func(j *JWT) {
		j.refreshTTL = duration
	}
}

func NewJWT(key string) JWTer {
	return NewJWTWithOptions(key)
}

func NewJWTWithOptions(key string, opts ...func(*JWT)) JWTer {
	j := &JWT{
		key:        []byte(key),
		ttl:        1 * time.Hour,
		refreshTTL: 14 * 24 * time.Hour,
	}

	for _, f := range opts {
		f(j)
	}

	return j
}

func (j *JWT) CreateClaims(userID string, issuer string) Claims {
	now := time.Now()

	return Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    issuer,
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(j.ttl)),
			Subject:   userID,
		},
	}
}

func (j *JWT) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.key)
}

func (j *JWT) ParseToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return j.key, nil
	})
}

func (j *JWT) RefreshToken(token string) (string, error) {
	tok, err := j.ParseToken(token)
	if err != nil {
		return "", err
	}

	if payload, ok := tok.Claims.(*Claims); ok {
		refreshExpiresAt := payload.IssuedAt.Add(j.refreshTTL)
		if refreshExpiresAt.Before(time.Now()) || payload.IssuedAt.After(time.Now()) {
			return "", ErrRefreshTokenExpired
		}

		expiresAt := time.Now().Add(j.refreshTTL)
		if refreshExpiresAt.Before(expiresAt) {
			expiresAt = refreshExpiresAt
		}

		return j.CreateToken(Claims{
			RegisteredClaims: jwt.RegisteredClaims{
				ID:        uuid.New().String(),
				IssuedAt:  payload.IssuedAt,
				Issuer:    payload.Issuer,
				NotBefore: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(expiresAt),
				Subject:   payload.Subject,
			},
		})
	}

	return "", ErrTokenInvalid
}

func (j *JWT) Decode(token string) (*Claims, error) {
	payload := strings.Split(token, ".")
	if len(payload) != 3 {
		return nil, ErrTokenInvalid
	}

	decodeBytes, err := jwt.NewParser().DecodeSegment(payload[1])
	if err != nil {
		return nil, err
	}

	var claims Claims
	if err := json.Unmarshal(decodeBytes, &claims); err != nil {
		return nil, err
	}

	return &claims, nil
}

func (j *JWT) TTL() time.Duration {
	return j.ttl
}

func (j *JWT) RefreshTTL() time.Duration {
	return j.refreshTTL
}
