package cmd

import (
	"go-boilerplate/internal/app"
	"go-boilerplate/internal/pkg/database"
	"go-boilerplate/pkg/config"

	"github.com/urfave/cli/v2"
)

// Start Start Server
func Start() *cli.Command {
	return &cli.Command{
		Name:        "start",
		Usage:       "Start Server",
		Description: "Start Server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "./config.yml",
				Usage:   "Load configuration from `FILE`",
			},
		},
		Action: func(c *cli.Context) error {
			if err := bootstrap(c); err != nil {
				return err
			}

			// migrate
			database.Migrate()

			if err := app.Init(); err != nil {
				return err
			}

			srv := app.NewServer()
			if config.Server.HTTPS {
				return srv.StartTLS()
			}

			return srv.Start()
		},
	}
}
