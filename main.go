package main

import (
	"errors"
	"fmt"
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
			utils.LoadGlobalConfig(cli.FlagGlobalConfigVar)
			monitor.LoadMonitorConfig(cli.FlagMonitorConfigVar)
			aliyun.LoadAliyunConfig(cli.FlagAliyunConfigVar)

			var dbConf = utils.GlobalConfig.Database
			var statsDbConf = utils.GlobalConfig.StatsDatabase
			var err error

			if utils.AnyMatch("", dbConf.Host, dbConf.User, dbConf.Password, dbConf.DBName) || dbConf.Port == 0 || utils.AnyMatch("", statsDbConf.Host, statsDbConf.User, statsDbConf.Password, statsDbConf.DBName) || statsDbConf.Port == 0 {
				return errors.New("failed to connect to database due to invalid configuration")
			}

			utils.DB = &utils.Database{}
			utils.StatsDB = &utils.Database{}

			err, utils.DB.Conn = utils.Load(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DBName))
			if err != nil {
				return err
			}
			err, utils.StatsDB.Conn = utils.Load(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				statsDbConf.User, statsDbConf.Password, statsDbConf.Host, statsDbConf.Port, statsDbConf.DBName))
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
