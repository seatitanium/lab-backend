package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errors"
	"seatimc/backend/utils"
)

// 动态获取指定 regionId 下的 instanceId 实例的实时信息
func DescribeInstance(instanceId string, regionId string) (*aliyun.InstanceDescriptionRetrieved, *errors.CustomErr) {
	client, customErr := aliyun.CreateEcsClient()
	if customErr != nil {
		return nil, customErr
	}

	res, err := client.DescribeInstances(&ecs.DescribeInstancesRequest{
		RegionId:    tea.String(regionId),
		InstanceIds: tea.String("['" + instanceId + "']"),
	})

	if err != nil {
		return nil, errors.AliyunError(err)
	}

	var result aliyun.InstanceDescriptionRetrieved

	result.Exist = false

	for _, inst := range res.Body.Instances.Instance {
		result.Status = tea.StringValue(inst.Status)
		result.PublicIpAddress = tea.StringSliceValue(inst.PublicIpAddress.IpAddress)
		parsedTime, err := utils.ParseTimeRFC3339Aliyun(tea.StringValue(inst.CreationTime))
		if err != nil {
			return nil, errors.ServerError(err)
		}
		result.CreationTime = parsedTime
		result.Exist = true
	}

	return &result, nil
}
