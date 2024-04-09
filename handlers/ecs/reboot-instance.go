package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/utils"
)

func HandleRebootInstance() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request StopInstanceRequest

		if err := context.ShouldBindJSON(&request); err != nil {
			utils.RespondNG(context, "Invalid Request Body", "")
			return
		}

		err := ecs.RebootInstance(request.InstanceId, request.Force)

		if err != nil {
			utils.RespondNG(context, "RebootInstance failed: "+err.Error(), "重启实例时出现问题")
			return
		}

		utils.RespondOK(context, "Successfully started rebooting of instance (id="+request.InstanceId+")", "成功重启实例")
	}
}
