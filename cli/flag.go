package cli

import (
	cliv2 "github.com/urfave/cli/v2"
	"log"
)

var (
	FlagGlobalConfigVar  string
	FlagMonitorConfigVar string
	FlagAliyunConfigVar  string
)

var (
	FlagGlobalConfig = cliv2.PathFlag{
		Name:        "config",
		Aliases:     []string{"c"},
		Value:       "./config.yml",
		Usage:       "Configuration file path",
		Destination: &FlagGlobalConfigVar,
	}

	FlagMonitorConfig = cliv2.PathFlag{
		Name:        "monitor-config",
		Aliases:     []string{"cm"},
		Value:       "./monitor/monitor.yml",
		Usage:       "Monitor configuration file path",
		Destination: &FlagMonitorConfigVar,
	}

	FlagAliyunConfig = cliv2.PathFlag{
		Name:        "aliyun-config",
		Aliases:     []string{"ca"},
		Value:       "./ecs/aconfig.yml",
		Usage:       "Aliyun configuration file path",
		Destination: &FlagAliyunConfigVar,
	}

	FlagHelp = cliv2.BoolFlag{
		Name:               "help",
		Aliases:            []string{"h"},
		Usage:              "Show help",
		DisableDefaultText: true,
		Action: func(ctx *cliv2.Context, b bool) error {
			err := cliv2.ShowAppHelp(ctx)
			if err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}
)
