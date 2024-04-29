package main

import (
	cliv2 "github.com/urfave/cli/v2"
	"log"
	"os"
	"seatimc/backend/aliyun"
	"seatimc/backend/cli"
	"seatimc/backend/monitor"
	"seatimc/backend/utils"
)

func main() {
	app := &cliv2.App{
		Name:  "tisea",
		Usage: "Take control of the backend.",
		Commands: []*cliv2.Command{
			&cli.CommandRun,
			&cli.CommandMonitor,
			&cli.CommandInit,
			&cli.CommandHelp,
		},
		Flags: []cliv2.Flag{
			&cli.FlagGlobalConfig,
			&cli.FlagMonitorConfig,
			&cli.FlagAliyunConfig,
			&cli.FlagHelp,
		},
		Before: func(ctx *cliv2.Context) error {
			utils.GlobalConfig.Load(cli.FlagGlobalConfigVar)
			monitor.MonitorConfig.Load(cli.FlagMonitorConfigVar)
			aliyun.AliyunConfig.Load(cli.FlagAliyunConfigVar)
			err := utils.DB.Load()
			if err != nil {
				return err
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
