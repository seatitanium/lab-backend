package ecs

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"seatimc/backend/ecs"
	"seatimc/backend/utils"
)

func HandleDeleteInstance(db *sqlx.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var request StopInstanceRequest

		if err := context.ShouldBindJSON(&request); err != nil {
			utils.RespondNG(context, "Invalid Request Body", "")
			return
		}

		err := utils.WriteManualEcsRecord(context, request.InstanceId, "delete", request.Force)

		if err != nil {
			utils.RespondNG(context, "Cannot write manual 'delete' record: "+err.Error(), "无法写入操作记录")
			return
		}

		err = ecs.DeleteInstance(request.InstanceId, request.Force)

		if err != nil {
			utils.RespondNG(context, "DeleteInstance failed: "+err.Error(), "删除实例时出现问题")
			return
		}

		utils.RespondOK(context, "Successfully deleted instance (id="+request.InstanceId+")", "成功删除实例")
	}
}
