package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"log"
)

// 获取一个由所有当前 region 下可用区的 id 构成的简单字符串数组
func DescribeZoneIds() []string {
	zones := DescribeZones()
	var result []string
	for _, z := range zones {
		result = append(result, z.ZoneId)
	}
	return result
}

// 获取 aconfig.yml 中设定的区域的所有可用区信息，返回一个 AvailableZone 数组
func DescribeZones() []AvailableZone {
	client, err := CreateClient()

	if err != nil {
		return nil
	}

	request := &ecs.DescribeZonesRequest{
		RegionId: tea.String(Conf().PrimaryRegionId),
	}

	resp, err := client.DescribeZones(request)

	if err != nil {
		log.Println("Error in DescribeZones: " + err.Error())
	}

	var result []AvailableZone

	zones := resp.Body.Zones.Zone
	for _, zone := range zones {
		var _availableInstanceTypes []string
		for _, typ := range zone.AvailableInstanceTypes.InstanceTypes {
			_availableInstanceTypes = append(_availableInstanceTypes, tea.StringValue(typ))
		}
		result = append(result, AvailableZone{
			AvailableInstanceTypes: _availableInstanceTypes,
			ZoneId:                 tea.StringValue(zone.ZoneId),
			ZoneLocalName:          tea.StringValue(zone.LocalName),
		})
	}

	return result
}
