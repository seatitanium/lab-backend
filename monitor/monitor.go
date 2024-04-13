package monitor

import (
	"github.com/jmoiron/sqlx"
	"time"
)

func Run(db *sqlx.DB) {
	runStoppedInstanceEnd := make(<-chan bool)

	go RunStoppedInstanceMonitor(db, time.Second*2, time.Hour*1, runStoppedInstanceEnd)
}
