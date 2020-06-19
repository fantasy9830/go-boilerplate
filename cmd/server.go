package cmd

import (
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/models"
	"go-boilerplate/pkg/mqtt"
	"go-boilerplate/pkg/server"
	"go-boilerplate/pkg/websocket"

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
			models.Migrate()

			// init mqtt
			if err := mqtt.Init(); err != nil {
				return err
			}

			// init websocket
			if err := websocket.Init(); err != nil {
				return err
			}

			srv := server.NweServer()
			if config.Server.HTTPS {
				return srv.StartTLS()
			}

			return srv.Start()
		},
	}
}
