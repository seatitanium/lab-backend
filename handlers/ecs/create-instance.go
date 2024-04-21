package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/ecs"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleCreateInstance(ctx *gin.Context) *errHandler.CustomErr {
	hasActive, customErr := utils.HasActiveInstance()
	if customErr != nil {
		return customErr
	}

	if hasActive == true {
		return errHandler.OperateNotApplied()
	}

	conf := ecs.AliyunConfig
	created, customErr := ecs.CreateInstance(conf)
	if customErr != nil {
		return customErr
	}

	customErr = utils.SaveNewActiveInstance(created, conf.PrimaryRegionId, conf.Using.InstanceType)
	if customErr != nil {
		return customErr
	}

	customErr = utils.WriteManualEcsRecord(ctx, "", "create", false)
	if customErr != nil {
		return customErr
	}

	middleware.RespSuccess(ctx)
	return nil
}
