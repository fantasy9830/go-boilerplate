package cmd

import (
	"go-boilerplate/pkg/models"

	"github.com/urfave/cli/v2"
)

// Migrate Run the database migrations
func Migrate() *cli.Command {
	return &cli.Command{
		Name:        "migrate",
		Usage:       "Run the database migrations",
		Description: "Run the database migrations",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "./config.yml",
				Usage:   "Load configuration from `FILE`",
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Create the migration repository",
				Action: func(c *cli.Context) error {
					if err := bootstrap(c); err != nil {
						return err
					}

					// migrate
					models.Migrate()

					return nil
				},
			},
			{
				Name:  "refresh",
				Usage: "Drop all tables and re-run all migrations",
				Action: func(c *cli.Context) error {
					if err := bootstrap(c); err != nil {
						return err
					}

					// reverse
					models.Reverse()

					// migrate
					models.Migrate()

					return nil
				},
			},
			{
				Name:  "reverse",
				Usage: "Rollback all database migrations",
				Action: func(c *cli.Context) error {
					if err := bootstrap(c); err != nil {
						return err
					}

					// reverse
					models.Reverse()

					return nil
				},
			},
		},
	}
}
