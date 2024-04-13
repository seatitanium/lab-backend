package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func WriteManualEcsRecord(db *sqlx.DB, context *gin.Context, instanceId string, actionType string, force bool) error {
	token := context.GetHeader("JWT")
	payload := ExtractJWTPayload(token)

	if payload == nil {
		return fmt.Errorf("JWT_TOKEN")
	}

	if force {
		actionType += "_force"
	}

	_, err := DbExec(db, "INSERT INTO `seati_ecs_record` (`instance_id`, `action_type`, `by_username`) VALUES (?, ?, ?)", instanceId, actionType, payload.Username)

	return err
}

func WriteAutomatedEcsRecord(db *sqlx.DB, instanceId string, actionType string, force bool) error {

	if force {
		actionType += "_force"
	}

	_, err := DbExec(db, "INSERT INTO `seati_ecs_record` (`instance_id`, `action_type`, `automated`) VALUES (?, ?, ?)", instanceId, actionType, 1)

	return err
}
