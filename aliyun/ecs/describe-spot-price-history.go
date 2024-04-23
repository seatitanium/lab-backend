package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"log"
	"seatimc/backend/aliyun"
	"seatimc/backend/utils"
	"time"
)

// 按照 aconfig.yml 中的配置获取参数中给定的 zoneId 对应的可用区的竞价数值历史，返回一个 SpotPriceHistory 数组
//   - 注：startTime 和 endTime 采用 yyyy-MM-ddTHH:mm:ssZ 格式，时区 UTC+0
func DescribeSpotPriceHistory(zoneId string, startTime string, endTime string) []aliyun.SpotPriceHistory {
	client, customErr := aliyun.CreateClient()

	if customErr != nil {
		return nil
	}

	conf := aliyun.AliyunConfig
	request := &ecs.DescribeSpotPriceHistoryRequest{
		RegionId:     tea.String(conf.PrimaryRegionId),
		NetworkType:  tea.String(conf.Using.NetworkType),
		IoOptimized:  tea.String(aliyun.GetIoOptimized(conf.Using.IoOptimized)),
		InstanceType: tea.String(conf.Using.InstanceType),
		ZoneId:       tea.String(zoneId),
		SpotDuration: tea.Int32(conf.Using.SpotDuration),
		OSType:       tea.String(conf.Using.OSType),
		StartTime:    tea.String(startTime),
		EndTime:      tea.String(endTime),
	}

	resp, err := client.DescribeSpotPriceHistory(request)

	if err != nil {
		log.Println("Error in DescribeSpotPriceHistory: " + err.Error())
	}

	var result []aliyun.SpotPriceHistory

	historys := resp.Body.SpotPrices.SpotPriceType
	for _, history := range historys {
		parsedTime, err := utils.ParseTime(tea.StringValue(history.Timestamp))
		if err != nil {
			log.Println("Warning in DescribeSpotPriceHistory: Cannot parse RFC3339 / ISO8601 time string in response. Replacing with timestamp 0.")
			parsedTime = time.UnixMilli(0)
		}
		result = append(result, aliyun.SpotPriceHistory{
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
