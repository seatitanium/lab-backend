package utils

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errHandler"
)

func WriteManualEcsRecord(ctx *gin.Context, instanceId string, actionType string, force bool) *errHandler.CustomErr {
	conn := GetDBConn()
	token := ctx.GetHeader("JWT")
	payload := ExtractJWTPayload(token)

	if payload == nil {
		return errHandler.UnAuth()
	}

	if force {
		actionType += "_force"
	}

	result := conn.Create(&EcsActions{
		InstanceId: instanceId,
		ActionType: actionType,
		ByUsername: payload.Username,
	})

	return errHandler.DbError(result.Error)
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

	return errHandler.DbError(result.Error)
}
