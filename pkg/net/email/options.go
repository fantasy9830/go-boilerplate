package email

import (
	"crypto/tls"
	"net/mail"

	"gopkg.in/gomail.v2"
)

type Options func(*Email)

func WithSMTP(host string, port int, username, password string, insecureSkipVerify bool) Options {
	return func(e *Email) {
		e.dialer = gomail.NewDialer(host, port, username, password)
		e.dialer.TLSConfig = &tls.Config{InsecureSkipVerify: insecureSkipVerify}
	}
}

func WithFrom(name, address string) Options {
	return func(e *Email) {
		e.from = mail.Address{
			Name:    name,
			Address: address,
		}
	}
}
