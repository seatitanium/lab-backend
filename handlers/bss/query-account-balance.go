package bss

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/aliyun/bss"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
)

func HandleQueryAccountBalance(ctx *gin.Context) *errHandler.CustomErr {
	result, customErr := bss.QueryAccountBalance()

	if customErr != nil {
		return customErr
	}

	middleware.RespSuccess(ctx, result)

	return nil
}
