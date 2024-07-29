package bss

import (
	"seatimc/backend/utils"
	"time"
)

var EcsConsumptionBefore202407 float32 = 1015.549999
var OSSConsumptionBefore202407 float32 = 148.18
var YundiskConsumptionBefore202407 float32 = 93.95

func GetConsumption(productCode string) float32 {
	startCycle := "2024-07"
	startCycleDate, err := utils.ParseTimeRFC3339(startCycle + "-01T00:00:00Z")

	if err != nil {
		return 0
	}

	var result float32 = 0

	switch productCode {
	case "ecs":
		result = EcsConsumptionBefore202407
	case "oss":
		result = OSSConsumptionBefore202407
	case "yundisk":
		result = YundiskConsumptionBefore202407
	default:
		return 0
	}

	now := time.Now()

	nowDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.FixedZone("UTC+8", 8*60*60))

	var cycle string

	for nowDay.After(startCycleDate) || nowDay.Equal(startCycleDate) {
		cycle = startCycleDate.Format("2006-01")
		resp, err := QueryBill(cycle, cycle, productCode, "PayAsYouGo")

		if err != nil {
			return 0
		}

		for _, b := range resp {
			result += b.CashAmount
		}

		startCycleDate = startCycleDate.AddDate(0, 1, 0)
	}

	return result
}
