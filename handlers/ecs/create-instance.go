package ecs

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"seatimc/backend/ecs"
	"seatimc/backend/utils"
)

func HandleCreateInstance(db *sqlx.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		hasActive, err := utils.HasActiveInstance(db)

		if err != nil {
			utils.RespondNG(context, "Unable to determine active instance: "+err.Error(), "获取活跃实例时出现问题")
			return
		}

		if hasActive == true {
			utils.RespondNG(context, "An active instance already exists.", "活跃实例已存在，不可重复创建")
			return
		}

		conf := ecs.AConf()
		created, err := ecs.CreateInstance(conf)

		if err != nil {
			utils.RespondNG(context, "Unable to create instance: "+err.Error(), "创建实例时出现问题")
			return
		}

		err = utils.SaveNewActiveInstance(db, created, conf.PrimaryRegionId, conf.Using.InstanceType)

		if err != nil {
			utils.RespondNG(context, "Unable to insert instance info into database: "+err.Error(), "无法保存实例")
		}

		utils.RespondOK(context, "Successfully created a new instance.", "成功创建新的实例")
	}
}
