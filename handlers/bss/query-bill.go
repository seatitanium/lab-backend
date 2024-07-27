package bss

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/aliyun"
	"seatimc/backend/aliyun/bss"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"time"
)

func HandleQueryBill(ctx *gin.Context) *errors.CustomErr {
	cycle := ctx.DefaultQuery("cycle", "")
	productCode := ctx.DefaultQuery("productCode", "")
	subscriptionType := ctx.DefaultQuery("subscriptionType", "")

	if productCode == "" || subscriptionType == "" {
		return errors.WrongParam()
	}

	var startCycle string
	var endCycle string
	var result []aliyun.Bill
	var customErr *errors.CustomErr

	if cycle == "" {
		startCycle = ctx.DefaultQuery("startCycle", "")
		endCycle = ctx.DefaultQuery("endCycle", "")

		if startCycle == "" {
			startCycle = "2021-01"
		}

		if endCycle == "" {
			endCycle = time.Now().Format("2006-01")
		}

		result, customErr = bss.QueryBill(startCycle, endCycle, productCode, subscriptionType)
	} else {
		result, customErr = bss.QueryBill(cycle, cycle, productCode, subscriptionType)
	}

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, result)

	return nil
}
