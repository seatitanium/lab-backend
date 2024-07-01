package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleDeleteInstance(ctx *gin.Context) *errors.CustomErr {
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

	customErr = ecs.DeleteInstance(request.InstanceId, request.Force)
	if customErr != nil {
		return customErr
	}

	customErr = utils.WriteManualEcsRecord(ctx, request.InstanceId, "delete", request.Force)
	if customErr != nil {
		return customErr
	}

	customErr = utils.SetActive(request.InstanceId, false)
	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx)
	return nil
}
