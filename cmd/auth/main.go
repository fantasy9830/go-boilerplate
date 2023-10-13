package main

import (
	"go-boilerplate/internal/auth/controller/http"
	"go-boilerplate/internal/auth/migration"
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/logger"
	"go-boilerplate/pkg/version"
	"log/slog"
	"os"
	"time"

	"github.com/fantasy9830/go-graceful"
	"github.com/urfave/cli/v2"
)

// @title Auth API
// @version 1.0
// @description This is an api for Auth Service.
// @BasePath /
func main() {
	app := &cli.App{
		Name:     "auth",
		Usage:    "Start the auth service",
		Version:  version.PrintCLIVersion(),
		Compiled: time.Now(),
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Usage:       "Debug mode",
				EnvVars:     []string{"GO_DEBUG"},
				Destination: &config.App.Debug,
			},
			&cli.StringFlag{
				Name:        "key",
				Usage:       "Secret key",
				EnvVars:     []string{"GO_SECRET_KEY"},
				Destination: &config.App.Key,
			},

			// postgres
			&cli.StringFlag{
				Name:        "postgres-host",
				Usage:       "postgres host",
				EnvVars:     []string{"GO_POSTGRES_HOST"},
				Destination: &config.Postgres.Host,
			},
			&cli.IntFlag{
				Name:        "postgres-port",
				Usage:       "postgres port",
				EnvVars:     []string{"GO_POSTGRES_PORT"},
				Destination: &config.Postgres.Port,
			},
			&cli.StringFlag{
				Name:        "postgres-username",
				Usage:       "postgres username",
				EnvVars:     []string{"GO_POSTGRES_USER"},
				Destination: &config.Postgres.Username,
			},
			&cli.StringFlag{
				Name:        "postgres-password",
				Usage:       "postgres password",
				EnvVars:     []string{"GO_POSTGRES_PASSWORD"},
				Destination: &config.Postgres.Password,
			},
			&cli.StringFlag{
				Name:        "postgres-dbname",
				Usage:       "postgres dbname",
				EnvVars:     []string{"GO_POSTGRES_DB"},
				Destination: &config.Postgres.DBName,
			},

			// redis
			&cli.StringFlag{
				Name:        "redis-type",
				Usage:       "Redis type",
				EnvVars:     []string{"GO_REDIS_TYPE"},
				Destination: &config.Redis.Type,
			},
			&cli.StringFlag{
				Name:        "redis-host",
				Usage:       "Redis host",
				EnvVars:     []string{"GO_REDIS_HOST"},
				Destination: &config.Redis.Host,
			},
			&cli.StringFlag{
				Name:        "redis-password",
				Usage:       "Redis password",
				EnvVars:     []string{"GO_REDIS_PASSWORD"},
				Destination: &config.Redis.Password,
			},

			// Mail
			&cli.StringFlag{
				Name:        "mail-smtp-host",
				Usage:       "smtp host",
				EnvVars:     []string{"GO_MAIL_SMTP_HOST"},
				Destination: &config.Mail.SMTP.Host,
			},
			&cli.IntFlag{
				Name:        "mail-smtp-port",
				Usage:       "smtp port",
				EnvVars:     []string{"GO_MAIL_SMTP_PORT"},
				Destination: &config.Mail.SMTP.Port,
			},
			&cli.StringFlag{
				Name:        "mail-smtp-username",
				Usage:       "smtp username",
				EnvVars:     []string{"GO_MAIL_SMTP_USERNAME"},
				Destination: &config.Mail.SMTP.Username,
			},
			&cli.StringFlag{
				Name:        "mail-smtp-password",
				Usage:       "smtp password",
				EnvVars:     []string{"GO_MAIL_SMTP_PASSWORD"},
				Destination: &config.Mail.SMTP.Password,
			},
			&cli.StringFlag{
				Name:        "mail-from-name",
				Usage:       "sender name",
				EnvVars:     []string{"GO_MAIL_FROM_NAME"},
				Destination: &config.Mail.From.Name,
			},
			&cli.StringFlag{
				Name:        "mail-from-address",
				Usage:       "sender address",
				EnvVars:     []string{"GO_MAIL_FROM_ADDRESS"},
				Destination: &config.Mail.From.Address,
			},
		},
		Before: func(ctx *cli.Context) error {
			logger.SetupLogger()
			if err := migration.Init(ctx.Context); err != nil {
				return err
			}

			return nil
		},
		Action: func(c *cli.Context) error {
			m := graceful.GetManager()

			http.Init()

			<-m.Done()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
