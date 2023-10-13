package auth_test

import (
	"go-boilerplate/pkg/auth"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JWTSuit struct {
	suite.Suite
	SecretKey string
}

func TestJWT(t *testing.T) {
	suite.Run(t, new(JWTSuit))
}

func (s *JWTSuit) SetupSuite() {
	s.SecretKey = "test"
}

func (s *JWTSuit) TestCreateClaims() {
	jwt := auth.NewJWT(s.SecretKey)
	claims := jwt.CreateClaims("1", "test_issuer")

	assert.Equal(s.T(), "1", claims.Subject)
	assert.Equal(s.T(), "test_issuer", claims.Issuer)
	assert.Equal(s.T(), claims.IssuedAt, claims.NotBefore)
}

func (s *JWTSuit) TestParseToken() {
	jwt := auth.NewJWT(s.SecretKey)
	claims := jwt.CreateClaims("1", "test_issuer")
	accessToken, err := jwt.CreateToken(claims)
	assert.NoError(s.T(), err)

	token, err := jwt.ParseToken(accessToken)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), accessToken, token.Raw)
}

func (s *JWTSuit) TestDecode() {
	jwt := auth.NewJWT(s.SecretKey)
	originalClaims := jwt.CreateClaims("1", "test_issuer")
	accessToken, err := jwt.CreateToken(originalClaims)
	assert.NoError(s.T(), err)

	decodeClaims, err := jwt.Decode(accessToken)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), originalClaims.ID, decodeClaims.ID)
	assert.Equal(s.T(), originalClaims.Subject, decodeClaims.Subject)

	_, err = jwt.Decode("")
	assert.EqualError(s.T(), err, "token is invalid")
}

func (s *JWTSuit) TestWithOptions() {
	jwt := auth.NewJWTWithOptions(s.SecretKey, auth.WithTTL(1*time.Hour), auth.WithRefreshTTL(2*time.Hour))

	assert.Equal(s.T(), jwt.TTL(), 1*time.Hour)
	assert.Equal(s.T(), jwt.RefreshTTL(), 2*time.Hour)
}

func (s *JWTSuit) TestRefreshToken() {
	jwt := auth.NewJWT(s.SecretKey)
	claims := jwt.CreateClaims("1", "test_issuer")
	accessToken, err := jwt.CreateToken(claims)
	assert.NoError(s.T(), err)

	refreshToken, err := jwt.RefreshToken(accessToken)
	assert.NoError(s.T(), err)

	assert.NotEqual(s.T(), accessToken, refreshToken)

	accessClaims, err := jwt.Decode(accessToken)
	assert.NoError(s.T(), err)

	refreshClaims, err := jwt.Decode(refreshToken)
	assert.NoError(s.T(), err)

	assert.NotEqual(s.T(), accessClaims.ID, refreshClaims.ID)
	assert.Equal(s.T(), accessClaims.Issuer, refreshClaims.Issuer)
	assert.Equal(s.T(), accessClaims.IssuedAt, refreshClaims.IssuedAt)
	assert.Equal(s.T(), accessClaims.Subject, refreshClaims.Subject)
}

func (s *JWTSuit) TestRefreshError() {
	jwt := auth.NewJWTWithOptions(s.SecretKey, auth.WithRefreshTTL(0))
	claims := jwt.CreateClaims("1", "test_issuer")
	accessToken, err := jwt.CreateToken(claims)
	assert.NoError(s.T(), err)

	_, err = jwt.RefreshToken(accessToken)
	assert.EqualError(s.T(), err, "token has expired and can no longer be refreshed")

	_, err = jwt.RefreshToken("")
	assert.EqualError(s.T(), err, "token is malformed: token contains an invalid number of segments")
}
