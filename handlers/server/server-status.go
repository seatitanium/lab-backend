package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mcstatus-io/mcutil/v3"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
	"strconv"
	"time"
)

func HandleServerStatus(ctx *gin.Context) *errors.CustomErr {
	mcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ip := ctx.DefaultQuery("ip", "")
	port := ctx.DefaultQuery("port", "25565")

	if ip == "" {
		hasActiveInstance, customErr := utils.HasActiveInstance()

		if customErr != nil {
			return customErr
		}

		if hasActiveInstance == false {
			return errors.WrongParam()
		}

		activeInstance, customErr := utils.GetActiveInstance()

		if customErr != nil {
			return customErr
		}

		ip = activeInstance.Ip
	}

	portInt, err := strconv.Atoi(port)

	if err != nil {
		return errors.ServerError(err)
	}

	resp, err := mcutil.Status(mcCtx, ip, uint16(portInt))

	if resp == nil {
		return errors.Offline()
	}

	handlers.RespSuccess(ctx, resp)

	return nil
}
