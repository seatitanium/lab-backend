package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleGetInvocationResult(ctx *gin.Context) *errors.CustomErr {
	var invokeId string
	var customErr *errors.CustomErr

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
			return errors.TargetNotExist()
		}
	}

	res, customErr := ecs.DescribeInvocationResults(invokeId)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, res)

	return nil
}
