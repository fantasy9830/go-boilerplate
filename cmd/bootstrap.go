package cmd

import (
	"go-boilerplate/internal/pkg/database"
	"go-boilerplate/pkg/config"

	"github.com/urfave/cli/v2"
)

func bootstrap(c *cli.Context) error {
	// init config
	if err := config.Load(c.String("config")); err != nil {
		return err
	}

	// init database
	if err := database.Init(); err != nil {
		return err
	}

	return nil
}
