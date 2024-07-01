package ecs

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleGetDeployStatus(ctx *gin.Context) *errors.CustomErr {
	hasActiveInstance, customErr := utils.HasActiveInstance()

	if customErr != nil {
		return customErr
	}

	if hasActiveInstance == false {
		return errors.TargetNotExist()
	}

	activeInstance, customErr := utils.GetActiveInstance()

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, activeInstance.DeployStatus)

	return nil
}
