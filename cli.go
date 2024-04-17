package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"os/signal"
	"seatimc/backend/handlers"
	"seatimc/backend/monitor"
	"seatimc/backend/utils"
	"syscall"
	"time"
)

func Run() {
	log.Println("Starting 🌊Tisea Backend.")

	router := handlers.Router{Port: utils.GlobalConfig.BindPort}

	router.Init()
	router.Run()
}

func RunMonitor(monitorName string) {

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
