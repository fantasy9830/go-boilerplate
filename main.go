package main

import (
	"go-boilerplate/cmd"
	"go-boilerplate/pkg/config"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = config.App.Name
	app.Usage = config.App.Usage
	app.Version = config.App.Version
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		cmd.Start(),
		cmd.Migrate(),
		cmd.Seed(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
