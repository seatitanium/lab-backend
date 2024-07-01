package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleStartInstance(ctx *gin.Context) *errors.CustomErr {
	instanceId := ctx.Query("instanceId")
	var customErr *errors.CustomErr

	if instanceId == "" {
		instanceId, customErr = utils.GetActiveInstanceId()

		if customErr != nil {
			return customErr
		}
	}

	customErr = ecs.StartInstance(instanceId)
	if customErr != nil {
		return customErr
	}

	customErr = utils.WriteManualEcsRecord(ctx, instanceId, "start", false)
	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx)
	return nil
}
