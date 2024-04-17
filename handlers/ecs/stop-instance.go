package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleStopInstance(ctx *gin.Context) *errHandler.CustomErr {
	var request StopInstanceRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		return errHandler.WrongParam()
	}

	customErr := utils.WriteManualEcsRecord(ctx, request.InstanceId, "stop", request.Force)
	if customErr != nil {
		return customErr
	}

	customErr = ecs.StopInstance(request.InstanceId, request.Force)
	if customErr != nil {
		return customErr
	}

	middleware.RespSuccess(ctx)
	return nil
}
