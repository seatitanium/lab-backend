package server

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandlePeakOnlineHistory(ctx *gin.Context) *errors.CustomErr {
	maximum, customErr := utils.GetPeakOnlineHistory()

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, maximum)

	return nil
}
