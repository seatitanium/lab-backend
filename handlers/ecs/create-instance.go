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

	conf := aliyun.AliyunConfig

	zoneId, customErr := ecs.GetAvailableZoneId(conf.Using.InstanceType)
	targetInstanceType := ""

	if customErr != nil {
		return customErr
	}

	if zoneId == "" {
		log.Printf("Warn: primary instance type [%v] is not available across the region.\n", conf.Using.InstanceType)

		if conf.Using.AltInstanceType == "" {
			log.Println("Warn: no alternative instance type configured.")
			return errors.OperationNotApplied()
		}

		zoneId, customErr = ecs.GetAvailableZoneId(conf.Using.AltInstanceType)

		if zoneId == "" {
			log.Printf("Warn: alternative instance type [%v] is not available across the region.\n", conf.Using.AltInstanceType)
			return errors.OperationNotApplied()
		} else {
			targetInstanceType = conf.Using.AltInstanceType
		}
	} else {
		targetInstanceType = conf.Using.InstanceType
	}

	created, customErr := ecs.CreateInstance(targetInstanceType, zoneId, conf)

	if customErr != nil {
		return customErr
	}

	customErr = utils.SaveNewActiveInstance(created, conf.PrimaryRegionId, zoneId, targetInstanceType)
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
