package server

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
	"time"
)

func HandleOnlineHistory(ctx *gin.Context) *errors.CustomErr {
	startTime := ctx.DefaultQuery("start", "")
	endTime := ctx.DefaultQuery("end", "")

	if startTime == "" {
		return errors.WrongParam()
	}

	parsedStart, err := utils.ParseTimeRFC3339Milli(startTime)

	if err != nil {
		return errors.WrongParam()
	}

	var parsedEnd time.Time

	if endTime != "" {
		parsedEnd, err = utils.ParseTimeRFC3339Milli(endTime)

		if err != nil {
			return errors.WrongParam()
		}
	} else {
		parsedEnd = time.Now()
	}

	history, customErr := utils.GetOnlineHistory(parsedStart, parsedEnd)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, history)
	return nil
}
