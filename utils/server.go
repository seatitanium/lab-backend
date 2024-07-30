package utils

import (
	"context"
	"github.com/mcstatus-io/mcutil/v3"
	"github.com/mcstatus-io/mcutil/v3/options"
	"github.com/mcstatus-io/mcutil/v3/response"
	"time"
)

func GetServerStatus(ip string, port uint16) (*response.JavaStatus, error) {
	mcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return mcutil.Status(mcCtx, ip, port, options.JavaStatus{
		Timeout: 5 * time.Second,
	})
}
