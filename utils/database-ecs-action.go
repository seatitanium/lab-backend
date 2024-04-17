package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func WriteManualEcsRecord(context *gin.Context, instanceId string, actionType string, force bool) error {
	conn := GetDBConn()
	token := context.GetHeader("JWT")
	payload := ExtractJWTPayload(token)

	if payload == nil {
		return fmt.Errorf("JWT_TOKEN")
	}

	if force {
		actionType += "_force"
	}

	result := conn.Create(&EcsActions{
		InstanceId: instanceId,
		ActionType: actionType,
		ByUsername: payload.Username,
	})

	return result.Error
}

func WriteAutomatedEcsRecord(instanceId string, actionType string, force bool) error {
	conn := GetDBConn()
	if force {
		actionType += "_force"
	}

	result := conn.Create(&EcsActions{
		InstanceId: instanceId,
		ActionType: actionType,
		Automated:  true,
	})

	return result.Error
}
