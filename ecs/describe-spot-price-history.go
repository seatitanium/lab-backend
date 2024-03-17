package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"log"
	"time"
)

func DescribeSpotPriceHistory(zoneId string) []SpotPriceHistory {
	client, err := CreateClient()

	if err != nil {
		return nil
	}

	var ioOptimized string
	if AConf().Using.IoOptimized {
		ioOptimized = "optimized"
	} else {
		ioOptimized = "none"
	}

	request := &ecs.DescribeSpotPriceHistoryRequest{
		RegionId:     tea.String(AConf().PrimaryRegionId),
		NetworkType:  tea.String(AConf().Using.NetworkType),
		IoOptimized:  tea.String(ioOptimized),
		InstanceType: tea.String(AConf().Using.InstanceType),
		ZoneId:       tea.String(zoneId),
		SpotDuration: tea.Int32(AConf().Using.SpotDuration),
		OSType:       tea.String(AConf().Using.OSType),
	}

	resp, err := client.DescribeSpotPriceHistory(request)

	if err != nil {
		log.Println("Error in DescribeSpotPriceHistory: " + err.Error())
	}

	var result []SpotPriceHistory

	historys := resp.Body.SpotPrices.SpotPriceType
	for _, history := range historys {
		parsedTime, err := time.Parse(time.RFC3339, tea.StringValue(history.Timestamp))
		if err != nil {
			log.Println("Warning in DescribeSpotPriceHistory: Cannot parse RFC3339 / ISO8601 time string in response. Replacing with timestamp 0.")
			parsedTime = time.UnixMilli(0)
		}
		result = append(result, SpotPriceHistory{
			Price:        tea.Float32Value(history.SpotPrice),
			OriginPrice:  tea.Float32Value(history.OriginPrice),
			TimeISO8601:  tea.StringValue(history.Timestamp),
			Timestamp:    parsedTime.UnixMilli(),
			InstanceType: tea.StringValue(history.InstanceType),
			ZoneId:       tea.StringValue(history.ZoneId),
			NetworkType:  tea.StringValue(history.NetworkType),
		})
	}

	return result
}
