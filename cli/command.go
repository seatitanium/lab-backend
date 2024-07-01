package cli

import (
	cliv2 "github.com/urfave/cli/v2"
	"log"
	"seatimc/backend/common"
	"seatimc/backend/monitor"
	"seatimc/backend/utils"
)

var (
	CommandRun = cliv2.Command{
		Name:    "run",
		Aliases: []string{"boot"},
		Usage:   "Start the backend service",
		Action: func(ctx *cliv2.Context) error {
			log.Println("Starting 🌊Tisea Backend.")
			router := common.Router{Port: utils.GlobalConfig.BindPort}
			router.Init()
			router.Run()
			return nil
		},
	}

	CommandMonitor = cliv2.Command{
		Name:  "monitor",
		Usage: "Start one of the monitors of backend",
		Action: func(ctx *cliv2.Context) error {
			monitorName := ctx.Args().Get(0)
			monitor.RunMonitor(monitorName)
			return nil
		},
	}

	CommandInit = cliv2.Command{
		Name:  "init",
		Usage: "Init backend database",
		Action: func(ctx *cliv2.Context) error {
			if err := utils.InitDB(); err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}

	CommandHelp = cliv2.Command{
		Name:  "help",
		Usage: "Show help",
		Action: func(ctx *cliv2.Context) error {
			err := cliv2.ShowAppHelp(ctx)
			if err != nil {
				return err
			}
			return nil
		},
	}
)
