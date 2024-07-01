package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/aliyun"
	"seatimc/backend/ecs"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleCreateInstance(ctx *gin.Context) *errors.CustomErr {
	hasActive, customErr := utils.HasActiveInstance()
	if customErr != nil {
		return customErr
	}

	if hasActive == true {
		return errors.OperateNotApplied()
	}

	conf := aliyun.AliyunConfig
	created, customErr := ecs.CreateInstance(conf)
	if customErr != nil {
		return customErr
	}

	customErr = utils.SaveNewActiveInstance(created, conf.PrimaryRegionId, conf.Using.InstanceType)
	if customErr != nil {
		return customErr
	}

	customErr = utils.WriteManualEcsRecord(ctx, created.InstanceId, "create", false)
	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx)
	return nil
}
