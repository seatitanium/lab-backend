package bss

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/aliyun/bss"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
)

func HandleQueryConsumption(ctx *gin.Context) *errors.CustomErr {
	result := bss.GetEcsConsumption()

	handlers.RespSuccess(ctx, result)

	return nil
}
