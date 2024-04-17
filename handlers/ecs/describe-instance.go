package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleDescribeInstance(ctx *gin.Context) *errHandler.CustomErr {
	var request CommonInstanceRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		return errHandler.WrongParam()
	}

	activeInstance, customErr := utils.GetInstanceByInstanceId(request.InstanceId)
	if customErr != nil {
		return customErr
	}

	retrieved, customErr := ecs.DescribeInstance(activeInstance.InstanceId, activeInstance.RegionId)
	if customErr != nil {
		return customErr
	}

	ecsDesc := ecs.InstanceDescription{
		Local: ecs.InstanceDescriptionLocal{
			InstanceId:   activeInstance.InstanceId,
			RegionId:     activeInstance.RegionId,
			InstanceType: activeInstance.InstanceType,
		},
		Retrieved: *retrieved,
	}
	middleware.RespSuccess(ctx, ecsDesc)
	return nil
}
