package utils

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errHandler"
)

func WriteManualEcsRecord(ctx *gin.Context, instanceId string, actionType string, force bool) *errHandler.CustomErr {
	conn := GetDBConn()
	token := ctx.GetHeader("Token")
	payload, err := GetPayloadFromToken(token)

	if err != nil {
		return err
	}

	if force {
		actionType += "_force"
	}

	result := conn.Create(&EcsActions{
		InstanceId: instanceId,
		ActionType: actionType,
		ByUsername: payload.Username,
	})

	if result.Error != nil {
		return errHandler.DbError(result.Error)
	}

	return nil
}

func WriteAutomatedEcsRecord(instanceId string, actionType string, force bool) *errHandler.CustomErr {
	conn := GetDBConn()
	if force {
		actionType += "_force"
	}

	result := conn.Create(&EcsActions{
		InstanceId: instanceId,
		ActionType: actionType,
		Automated:  true,
	})

	if result.Error != nil {
		return errHandler.DbError(result.Error)
	}

	return nil
}
