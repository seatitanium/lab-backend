package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleStopInstance(ctx *gin.Context) *errors.CustomErr {
	var request StopInstanceRequest
	var customErr *errors.CustomErr

	request.InstanceId = ctx.Query("instanceId")
	request.Force = utils.IsTrue(ctx.Query("force"))

	if request.InstanceId == "" {
		request.InstanceId, customErr = utils.GetActiveInstanceId()

		if customErr != nil {
			return customErr
		}
	}

	customErr = ecs.StopInstance(request.InstanceId, request.Force)
	if customErr != nil {
		return customErr
	}

	customErr = utils.WriteManualEcsRecord(ctx, request.InstanceId, "stop", request.Force)
	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx)
	return nil
}
