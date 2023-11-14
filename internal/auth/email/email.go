package email

import (
	"go-boilerplate/internal/auth/config"
	"go-boilerplate/pkg/net/email"
	"net/mail"
	"sync"
)

var (
	once sync.Once
	e    email.Emailer
)

func Get() email.Emailer {
	once.Do(func() {
		e = email.NewMail(email.Config{
			Host:     config.Mail.SMTP.Host,
			Port:     config.Mail.SMTP.Port,
			Username: config.Mail.SMTP.Username,
			Password: config.Mail.SMTP.Password,
			Address: mail.Address{
				Name:    config.Mail.From.Name,
				Address: config.Mail.From.Address,
			},
		})
	})

	return e
}
