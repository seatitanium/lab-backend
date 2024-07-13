package server

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleOnlineHistory(ctx *gin.Context) *errors.CustomErr {
	startTime := ctx.DefaultQuery("start", "")
	endTime := ctx.DefaultQuery("end", "")

	if startTime == "" || endTime == "" {
		return errors.WrongParam()
	}

	parsedStart, err := utils.ParseTimeRFC3339Milli(startTime)

	if err != nil {
		return errors.WrongParam()
	}

	parsedEnd, err := utils.ParseTimeRFC3339Milli(endTime)

	if err != nil {
		return errors.WrongParam()
	}

	history, customErr := utils.GetOnlineHistory(parsedStart, parsedEnd)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, history)
	return nil
}
