package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/utils"
)

func HandleStartInstance() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request CommonInstanceRequest

		if err := context.ShouldBindJSON(&request); err != nil {
			utils.RespondNG(context, "Invalid Request Body", "")
			return
		}

		err := utils.WriteManualEcsRecord(context, request.InstanceId, "start", false)

		if err != nil {
			utils.RespondNG(context, "Cannot write manual 'start' record: "+err.Error(), "无法写入操作记录")
			return
		}

		err = ecs.StartInstance(request.InstanceId)

		if err != nil {
			utils.RespondNG(context, "StartInstance failed: "+err.Error(), "开启实例时出现问题")
			return
		}

		utils.RespondOK(context, "Successfully started instance (id="+request.InstanceId+")", "成功开启实例")
	}
}
