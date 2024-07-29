package bss

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/aliyun/bss"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
)

func HandleQueryConsumption(ctx *gin.Context) *errors.CustomErr {
	ecs := bss.GetConsumption("ecs")
	oss := bss.GetConsumption("oss")
	yundisk := bss.GetConsumption("yundisk")
	var result = struct {
		Ecs     float32 `json:"ecs"`
		Oss     float32 `json:"oss"`
		Yundisk float32 `json:"yundisk"`
		Sum     float32 `json:"sum"`
	}{
		Ecs:     ecs,
		Oss:     oss,
		Yundisk: yundisk,
		Sum:     ecs + oss + yundisk,
	}

	handlers.RespSuccess(ctx, result)

	return nil
}
