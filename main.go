package backend

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "TiseaBackend",
		Usage: "Take control of the backend through this CLI.",
		Commands: []*cli.Command{{
			Name:    "boot",
			Aliases: []string{"run"},
			Usage:   "Start the backend service",
			Action: func(context *cli.Context) error {
				Boot()
				return nil
			},
		}},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
