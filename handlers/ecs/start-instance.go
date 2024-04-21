package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleStartInstance(ctx *gin.Context) *errHandler.CustomErr {
	instanceId := ctx.Query("instanceId")

	if instanceId == "" {
		return errHandler.WrongParam()
	}

	customErr := utils.WriteManualEcsRecord(ctx, instanceId, "start", false)
	if customErr != nil {
		return customErr
	}

	customErr = ecs.StartInstance(instanceId)
	if customErr != nil {
		return customErr
	}

	middleware.RespSuccess(ctx)
	return nil
}
