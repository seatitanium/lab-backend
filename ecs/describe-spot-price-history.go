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

	conf := AConf()
	request := &ecs.DescribeSpotPriceHistoryRequest{
		RegionId:     tea.String(conf.PrimaryRegionId),
		NetworkType:  tea.String(conf.Using.NetworkType),
		IoOptimized:  tea.String(GetIoOptimized(conf.Using.IoOptimized)),
		InstanceType: tea.String(conf.Using.InstanceType),
		ZoneId:       tea.String(zoneId),
		SpotDuration: tea.Int32(conf.Using.SpotDuration),
		OSType:       tea.String(conf.Using.OSType),
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
