package ecs

import (
	client2 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errors"
)

func GetAvailableZoneId(instanceType string) (string, *errors.CustomErr) {
	client, customErr := aliyun.CreateEcsClient()

	if customErr != nil {
		return "", customErr
	}

	resp, err := client.DescribeAvailableResource(&client2.DescribeAvailableResourceRequest{
		RegionId:            tea.String(aliyun.AliyunConfig.PrimaryRegionId),
		InstanceChargeType:  tea.String(aliyun.AliyunConfig.Using.InstanceChargeType),
		SpotStrategy:        tea.String(aliyun.AliyunConfig.Using.SpotStrategy),
		DestinationResource: tea.String("InstanceType"),
		InstanceType:        tea.String(instanceType),
	})

	if err != nil {
		return "", errors.AliyunError(err)
	}

	zoneId := ""

Main:
	for _, z := range resp.Body.AvailableZones.AvailableZone {
		for _, r := range z.AvailableResources.AvailableResource {
			for _, rr := range r.SupportedResources.SupportedResource {
				if tea.StringValue(rr.Status) == "Available" {
					zoneId = tea.StringValue(z.ZoneId)
					break Main
				}
			}
		}
	}

	return zoneId, nil
}
