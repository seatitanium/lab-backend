package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleDescribeInstance(ctx *gin.Context) *errHandler.CustomErr {
	instanceId := ctx.Query("instanceId")

	if instanceId == "" {
		return errHandler.WrongParam()
	}

	hasActiveInstance, customErr := utils.HasActiveInstance()

	if customErr != nil {
		return customErr
	}

	if hasActiveInstance == false {
		return errHandler.ResNotExist()
	}

	activeInstance, customErr := utils.GetInstanceByInstanceId(instanceId)
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
