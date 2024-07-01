package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/aliyun"
	"seatimc/backend/ecs"
	"seatimc/backend/errors"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
	"time"
)

func HandleDescribeInstance(ctx *gin.Context) *errors.CustomErr {
	instanceId := ctx.Query("instanceId")
	var customErr *errors.CustomErr
	var hasActiveInstance bool

	if instanceId == "" {
		hasActiveInstance, customErr = utils.HasActiveInstance()

		if hasActiveInstance == false {
			middleware.RespSuccess(ctx, aliyun.InstanceDescription{
				Retrieved: aliyun.InstanceDescriptionRetrieved{
					Exist:           false,
					Status:          "",
					PublicIpAddress: nil,
					CreationTime:    time.Time{},
				},
				Local: aliyun.InstanceDescriptionLocal{
					InstanceId:   "",
					RegionId:     "",
					InstanceType: "",
				},
			})
			return nil
		}

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
