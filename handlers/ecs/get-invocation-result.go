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
	var customErr *errHandler.CustomErr

	invokeId = ctx.Query("invokeId")

	if invokeId == "" {

		activeInstanceId, customErr := utils.GetActiveInstanceId()

		if customErr != nil {
			return customErr
		}

		invokeId, customErr = utils.GetLastInvokeId(activeInstanceId)

		if customErr != nil {
			return customErr
		}

		if invokeId == "" {
			return errHandler.TargetNotExist()
		}
	}

	res, customErr := ecs.DescribeInvocationResults(invokeId)

	if customErr != nil {
		return customErr
	}

	middleware.RespSuccess(ctx, res)

	return nil
}
