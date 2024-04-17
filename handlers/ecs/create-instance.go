package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/utils"
)

func HandleCreateInstance() gin.HandlerFunc {
	return func(context *gin.Context) {
		hasActive, err := utils.HasActiveInstance()

		if err != nil {
			utils.RespondNG(context, "Unable to determine active instance: "+err.Error(), "获取活跃实例时出现问题")
			return
		}

		if hasActive == true {
			utils.RespondNG(context, "An active instance already exists.", "活跃实例已存在，不可重复创建")
			return
		}

		err = utils.WriteManualEcsRecord(context, "", "create", false)

		if err != nil {
			utils.RespondNG(context, "Cannot write manual 'create' record: "+err.Error(), "无法写入操作记录")
			return
		}

		conf := ecs.Conf()
		created, err := ecs.CreateInstance(conf)

		if err != nil {
			utils.RespondNG(context, "Unable to create instance: "+err.Error(), "创建实例时出现问题")
			return
		}

		err = utils.SaveNewActiveInstance(created, conf.PrimaryRegionId, conf.Using.InstanceType)

		if err != nil {
			utils.RespondNG(context, "Unable to insert instance info into database: "+err.Error(), "无法保存实例")
		}

		utils.RespondOK(context, "Successfully created a new instance.", "成功创建新的实例")
	}
}
