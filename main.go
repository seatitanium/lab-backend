package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"seatimc/backend/utils"
)

func main() {
	app := &cli.App{
		Name:  "tisea",
		Usage: "Take control of the backend.",
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"boot"},
				Usage:   "Start the backend service",
				Action: func(context *cli.Context) error {
					Run()
					return nil
				},
			},
			{
				Name:  "monitor",
				Usage: "Start one of the monitors of backend",
				Action: func(context *cli.Context) error {
					monitorName := context.Args().Get(0)
					RunMonitor(monitorName)
					return nil
				},
			},
		},
	}

	// TODO: 作为参数引入
	utils.GlobalConfig.Load("./config.yml")

	if err := utils.InitDB(); err != nil {
		log.Fatal(err)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
