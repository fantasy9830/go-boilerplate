package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"net/url"
)

type Signaturer interface {
	Signed(data string) string
	Check(hashString string, data string) bool
	SignedRoute(u *url.URL) string
}

type Signature struct {
	Key []byte
}

func NewSignature(key string) Signaturer {
	return &Signature{
		Key: []byte(key),
	}
}

// Signed Signed
func (s *Signature) Signed(data string) string {
	mac := hmac.New(sha256.New, s.Key)
	mac.Write([]byte(data))

	return fmt.Sprintf("%x", mac.Sum(nil))
}

// SignedRoute SignedRoute
func (s *Signature) SignedRoute(u *url.URL) string {
	return s.Signed(u.String())
}

// Check Check
func (s *Signature) Check(hashString string, data string) bool {
	signature := s.Signed(data)

	return hmac.Equal([]byte(hashString), []byte(signature))
}
