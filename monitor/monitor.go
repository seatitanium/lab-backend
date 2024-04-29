package monitor

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
			go RunInstanceStatusMonitor(time.Second*5, time.Hour*1, b)
			break
		}

	case "deploy-inst":
		{
			go RunDeployMonitor(time.Second*5, b)
			break
		}

	default:
		{
			log.Printf("Monitor of name \"%v\" doesn't exist.\n", monitorName)
		}
	}

	<-b
	log.Printf("\nStopping monitor \"%v\"\n", monitorName)
}
