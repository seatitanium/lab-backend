package bss

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/aliyun/bss"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
)

func HandleQueryAccountBalance(ctx *gin.Context) *errors.CustomErr {
	result, customErr := bss.QueryAccountBalance()

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, result)

	return nil
}
