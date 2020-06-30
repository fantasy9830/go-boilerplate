package cmd

import (
	"go-boilerplate/internal/pkg/database"

	"github.com/urfave/cli/v2"
)

// Seed Seed the database with records
func Seed() *cli.Command {
	return &cli.Command{
		Name:        "seed",
		Usage:       "Seed the database with records",
		Description: "Seed the database with records",
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

			database.Seed()

			return nil
		},
	}
}
