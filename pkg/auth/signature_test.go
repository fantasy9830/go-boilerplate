package auth_test

import (
	"go-boilerplate/pkg/auth"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SignatureSuit struct {
	suite.Suite
	SecretKey string
}

func TestSignature(t *testing.T) {
	suite.Run(t, new(SignatureSuit))
}

func (s *SignatureSuit) SetupSuite() {
	s.SecretKey = "test"
}

func (s *SignatureSuit) TestSignature() {
	signature := auth.NewSignature(s.SecretKey)

	signedString := signature.Signed("test")

	assert.Equal(s.T(), true, signature.Check(signedString, "test"))
	assert.Equal(s.T(), false, signature.Check(signedString, "test!"))

	u, err := url.Parse("http://localhost:8080?test=true")
	assert.NoError(s.T(), err)
	signedRouteString := signature.SignedRoute(u)
	assert.Equal(s.T(), true, signature.Check(signedRouteString, u.String()))
}
