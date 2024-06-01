package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/aliyun"
	"seatimc/backend/ecs"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleDescribeInstance(ctx *gin.Context) *errHandler.CustomErr {
	instanceId := ctx.Query("instanceId")
	var customErr *errHandler.CustomErr

	if instanceId == "" {
		instanceId, customErr = utils.GetActiveInstanceId()

		if customErr != nil {
			return customErr
		}
	}

	targetInstance, customErr := utils.GetInstanceByInstanceId(instanceId)
	if customErr != nil {
		return customErr
	}

	retrieved, customErr := ecs.DescribeInstance(targetInstance.InstanceId, targetInstance.RegionId)
	if customErr != nil {
		return customErr
	}

	ecsDesc := aliyun.InstanceDescription{
		Local: aliyun.InstanceDescriptionLocal{
			InstanceId:   targetInstance.InstanceId,
			RegionId:     targetInstance.RegionId,
			InstanceType: targetInstance.InstanceType,
		},
		Retrieved: *retrieved,
	}

	middleware.RespSuccess(ctx, ecsDesc)
	return nil
}
