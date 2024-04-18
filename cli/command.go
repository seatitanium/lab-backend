package cli

import (
	cliv2 "github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"seatimc/backend/handlers/common"
	"seatimc/backend/monitor"
	"seatimc/backend/utils"
	"syscall"
	"time"
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
			runMonitor(monitorName)
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

// FIXME: Move to monitor/monitor.go
func runMonitor(monitorName string) {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	b := make(chan bool)

	// 当接收到中止信号时，将 b 设置为 true
	go func() {
		<-c
		b <- true
	}()

	switch monitorName {
	case "stopped-inst":
		{
			go monitor.RunStoppedInstanceMonitor(time.Second, time.Hour, b)
			break
		}

	default:
		{
			log.Printf("Monitor of name \" %v \" doesn't exist.\n", monitorName)
		}
	}
}
