package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/utils"
)

// 动态获取指定 regionId 下的 instanceId 实例的实时信息
func DescribeInstance(instanceId string, regionId string) (*InstanceDescriptionRetrieved, error) {
	client, err := CreateClient()

	if err != nil {
		return nil, err
	}

	res, err := client.DescribeInstances(&ecs.DescribeInstancesRequest{
		RegionId:    tea.String(regionId),
		InstanceIds: tea.String("[" + instanceId + "]"),
	})

	if err != nil {
		return nil, err
	}

	var result InstanceDescriptionRetrieved

	for _, inst := range res.Body.Instances.Instance {
		result.Status = tea.StringValue(inst.Status)
		result.PublicIpAddress = tea.StringSliceValue(inst.PublicIpAddress.IpAddress)
		parsedTime, err := utils.ParseTime(tea.StringValue(inst.CreationTime))
		if err != nil {
			return nil, err
		}
		result.CreationTime = parsedTime
	}

	return &result, nil
}
