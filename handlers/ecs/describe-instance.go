package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/utils"
)

func HandleDescribeInstance() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request CommonInstanceRequest

		if err := context.ShouldBindJSON(&request); err != nil {
			utils.RespondNG(context, "Invalid Request Body", "")
			return
		}

		activeInstance, err := utils.GetInstanceByInstanceId(request.InstanceId)

		if err != nil {
			utils.RespondNG(context, "Unable to get instance from database: "+err.Error(), "获取活跃实例时出现问题")
			return
		}

		retrieved, err := ecs.DescribeInstance(activeInstance.InstanceId, activeInstance.RegionId)
		local := ecs.InstanceDescriptionLocal{
			InstanceId:   activeInstance.InstanceId,
			RegionId:     activeInstance.RegionId,
			InstanceType: activeInstance.InstanceType,
		}

		if err != nil {
			utils.RespondNG(context, "DescribeInstance failed: "+err.Error(), "获取实例信息时出现问题")
			return
		}

		utils.ReturnOK(context, "Successfully got instance info.", "成功获取到实例信息", ecs.InstanceDescription{
			Local:     local,
			Retrieved: *retrieved,
		})
	}
}
