package ecs

import (
	"github.com/gin-gonic/gin"
	"log"
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
		return errors.OperationNotApplied()
	}

	zoneId, customErr := ecs.GetAvailableZoneId()

	if customErr != nil {
		return customErr
	}

	if zoneId == "" {
		log.Println("Warn: target instance type is not available across the region.")
		return errors.OperationNotApplied()
	}

	conf := aliyun.AliyunConfig

	created, customErr := ecs.CreateInstance(zoneId, conf)

	if customErr != nil {
		return customErr
	}

	customErr = utils.SaveNewActiveInstance(created, conf.PrimaryRegionId, zoneId, conf.Using.InstanceType)
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
