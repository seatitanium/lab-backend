package ecs

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"seatimc/backend/ecs"
	"seatimc/backend/utils"
)

func HandleStopInstance(db *sqlx.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var request StopInstanceRequest

		if err := context.ShouldBindJSON(&request); err != nil {
			utils.RespondNG(context, "Invalid Request Body", "")
			return
		}

		err := utils.WriteManualEcsRecord(db, context, request.InstanceId, "stop", request.Force)

		if err != nil {
			utils.RespondNG(context, "Cannot write manual 'stop' record: "+err.Error(), "无法写入操作记录")
			return
		}

		err = ecs.StopInstance(request.InstanceId, request.Force)

		if err != nil {
			utils.RespondNG(context, "StopInstance failed: "+err.Error(), "停止实例时出现问题")
			return
		}

		utils.RespondOK(context, "Successfully stopped instance (id="+request.InstanceId+")", "成功停止实例")
	}
}
