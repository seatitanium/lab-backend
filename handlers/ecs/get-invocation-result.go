package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleGetInvocationResult(ctx *gin.Context) *errHandler.CustomErr {
	var invokeId string

	invokeId = ctx.Query("invokeId")

	if invokeId == "" {

		hasActiveInstance, customErr := utils.HasActiveInstance()

		if customErr != nil {
			return customErr
		}

		if !hasActiveInstance {
			return errHandler.NotFound()
		}

		activeInstance, customErr := utils.GetActiveInstance()

		if customErr != nil {
			return customErr
		}

		invokeId, customErr = utils.GetLastInvokeId(activeInstance.InstanceId)

		if customErr != nil {
			return customErr
		}
	}

	res, customErr := ecs.DescribeInvocationResults(invokeId)

	if customErr != nil {
		return customErr
	}

	middleware.RespSuccess(ctx, res)

	return nil
}
