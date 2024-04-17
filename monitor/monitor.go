package monitor

import (
	"time"
)

func Run() {
	runStoppedInstanceEnd := make(<-chan bool)

	go RunStoppedInstanceMonitor(time.Second*2, time.Hour*1, runStoppedInstanceEnd)
}
