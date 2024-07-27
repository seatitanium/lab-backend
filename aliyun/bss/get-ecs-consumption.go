package bss

import (
	"seatimc/backend/utils"
	"time"
)

var EcsConsumption202407 float32 = 1015.55

func GetEcsConsumption() float32 {
	startCycle := "2024-08"
	startCycleDate, err := utils.ParseTimeRFC3339(startCycle + "-01T00:00:00Z")

	if err != nil {
		return 0
	}

	result := EcsConsumption202407

	now := time.Now()

	nowDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.FixedZone("UTC+8", 8*60*60))

	for nowDay.After(startCycleDate) || nowDay.Equal(startCycleDate) {
		resp, err := QueryBill(startCycle, startCycle, "ecs", "PayAsYouGo")

		if err != nil {
			return 0
		}

		for _, b := range resp {
			result += b.CashAmount
		}
	}

	return result
}
