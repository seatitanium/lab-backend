package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleDeleteInstance(ctx *gin.Context) *errHandler.CustomErr {
	var request StopInstanceRequest

	request.InstanceId = ctx.Query("instanceId")
	request.Force = utils.IsTrue(ctx.Query("force"))

	if request.InstanceId == "" {
		return errHandler.WrongParam()
	}

	customErr := utils.WriteManualEcsRecord(ctx, request.InstanceId, "delete", request.Force)

	if customErr != nil {
		return customErr
	}

	customErr = ecs.DeleteInstance(request.InstanceId, request.Force)
	if customErr != nil {
		return customErr
	}

	middleware.RespSuccess(ctx)
	return nil
}
