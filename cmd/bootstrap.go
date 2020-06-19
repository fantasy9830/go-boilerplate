package cmd

import (
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/models"

	"github.com/urfave/cli/v2"
)

func bootstrap(c *cli.Context) error {
	// init config
	if err := config.Load(c.String("config")); err != nil {
		return err
	}

	// init model
	if err := models.Init(); err != nil {
		return err
	}

	return nil
}
