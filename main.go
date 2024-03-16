package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "tisea",
		Usage: "Take control of the backend.",
		Commands: []*cli.Command{{
			Name:    "run",
			Aliases: []string{"boot"},
			Usage:   "Start the backend service",
			Action: func(context *cli.Context) error {
				Run()
				return nil
			},
		}},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
