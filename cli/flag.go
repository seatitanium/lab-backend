package cli

import (
	cliv2 "github.com/urfave/cli/v2"
	"log"
)

var (
	FlagConfigVar string
)

var (
	FlagConfig = cliv2.PathFlag{
		Name:        "config",
		Aliases:     []string{"c"},
		Value:       "./config.yml",
		Usage:       "Configuration file path",
		Destination: &FlagConfigVar,
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
