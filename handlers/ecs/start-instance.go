package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleStartInstance(ctx *gin.Context) *errHandler.CustomErr {
	var request CommonInstanceRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		return errHandler.WrongParam()
	}

	customErr := utils.WriteManualEcsRecord(ctx, request.InstanceId, "start", false)
	if customErr != nil {
		return customErr
	}

	customErr = ecs.StartInstance(request.InstanceId)
	if customErr != nil {
		return customErr
	}

	middleware.RespSuccess(ctx)
	return nil
}
