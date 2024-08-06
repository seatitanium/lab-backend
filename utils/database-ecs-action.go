package utils

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
)

func WriteManualEcsRecord(ctx *gin.Context, instanceId string, actionType string, force bool) *errors.CustomErr {
	conn := GetDBConn()

	var payload *JWTPayload
	var customErr *errors.CustomErr

	if !VerifyServerSecretCtx(ctx) {
		payload, customErr = GetPayloadFromToken(ctx.GetHeader("Authorization"))

		if customErr != nil {
			return customErr
		}
	} else {
		payload = &JWTPayload{
			Username:  "server",
			UpdatedAt: "",
		}
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
		return errors.DbError(result.Error)
	}

	return nil
}

func WriteAutomatedEcsRecord(instanceId string, actionType string, force bool) *errors.CustomErr {
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
		return errors.DbError(result.Error)
	}

	return nil
}
