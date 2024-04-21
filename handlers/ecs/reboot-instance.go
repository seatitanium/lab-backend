package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleRebootInstance(ctx *gin.Context) *errHandler.CustomErr {
	var request StopInstanceRequest

	request.InstanceId = ctx.Query("instanceId")
	request.Force = utils.IsTrue(ctx.Query("force"))

	if request.InstanceId == "" {
		return errHandler.WrongParam()
	}

	customErr := ecs.RebootInstance(request.InstanceId, request.Force)
	if customErr != nil {
		return customErr
	}

	customErr = utils.WriteManualEcsRecord(ctx, request.InstanceId, "reboot", request.Force)
	if customErr != nil {
		return customErr
	}

	middleware.RespSuccess(ctx)
	return nil
}
