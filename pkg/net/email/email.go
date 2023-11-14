package email

import (
	"context"
	"crypto/tls"
	"log/slog"
	"net/mail"

	"github.com/fantasy9830/go-retry"
	"gopkg.in/gomail.v2"
)

type Emailer interface {
	To(to []mail.Address) Emailer
	Cc(cc []mail.Address) Emailer
	Subject(subject string) Emailer
	Content(content string) Emailer
	Send() error
}

type Message struct {
	to      []string
	cc      []string
	subject string
	content string
}

type Email struct {
	dialer  *gomail.Dialer
	config  Config
	message Message
}

type Config struct {
	Host               string
	Port               int
	Username           string
	Password           string
	Address            mail.Address
	InsecureSkipVerify bool
}

func NewMail(cfg Config) Emailer {
	e := &Email{
		config: cfg,
		dialer: gomail.NewDialer(
			cfg.Host,
			cfg.Port,
			cfg.Username,
			cfg.Password,
		),
	}

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	e.dialer.TLSConfig = &tls.Config{InsecureSkipVerify: cfg.InsecureSkipVerify}

	return e
}

func (e *Email) To(to []mail.Address) Emailer {
	if len(to) > 0 {
		e.message.to = make([]string, len(to))
		for i := 0; i < len(to); i++ {
			e.message.to[i] = to[i].String()
		}
	}

	return e
}

func (e *Email) Cc(cc []mail.Address) Emailer {
	if len(cc) > 0 {
		e.message.cc = make([]string, len(cc))
		for i := 0; i < len(cc); i++ {
			e.message.cc[i] = cc[i].String()
		}
	}

	return e
}

func (e *Email) Subject(subject string) Emailer {
	e.message.subject = subject

	return e
}

func (e *Email) Content(content string) Emailer {
	e.message.content = content

	return e
}

func (e *Email) Send() error {
	m := gomail.NewMessage()
	// defer m.Reset()

	// From
	m.SetHeader("From", e.config.Address.String())

	// To
	m.SetHeader("To", e.message.to...)

	// Cc
	if len(e.message.cc) > 0 {
		m.SetHeader("Cc", e.message.cc...)
	}

	// Subject
	m.SetHeader("Subject", e.message.subject)

	// Content
	m.SetBody("text/html", e.message.content)

	err := retry.Do(func(ctx context.Context) error {
		return e.dialer.DialAndSend(m)
	})
	if err != nil {
		slog.Error(err.Error())
	}

	return err
}
