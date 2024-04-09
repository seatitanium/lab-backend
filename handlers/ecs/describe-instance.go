package ecs

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"seatimc/backend/ecs"
	"seatimc/backend/utils"
)

func HandleDescribeInstance(db *sqlx.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		hasActive, err := utils.HasActiveInstance(db)

		if err != nil {
			utils.RespondNG(context, "Unable to determine active instance: "+err.Error(), "获取活跃实例时出现问题")
			return
		}

		if hasActive == false {
			utils.ReturnOK(context, "No active instance record", "暂无活跃实例", nil)
			return
		}

		activeInstance, err := utils.GetActiveInstance(db)

		if err != nil {
			utils.RespondNG(context, "Unable to get instance from database: "+err.Error(), "获取活跃实例时出现问题")
			return
		}

		res, err := ecs.DescribeInstance(activeInstance.InstanceId, activeInstance.RegionId)

		if err != nil {
			utils.RespondNG(context, "Unable to DescribeInstance (aliyun): "+err.Error(), "调用 API 过程出现问题")
			return
		}

		utils.ReturnOK(context, "Successfully got instance info.", "成功获取到实例信息", res)
	}
}
