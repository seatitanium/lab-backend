package server

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
	"strconv"
)

func HandleServerStatus(ctx *gin.Context) *errors.CustomErr {
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

	resp, err := utils.GetServerStatus(ip, uint16(portInt))

	if err != nil {
		return errors.ServerError(err)
	}

	if resp == nil {
		return errors.Offline()
	}

	handlers.RespSuccess(ctx, resp)

	return nil
}
