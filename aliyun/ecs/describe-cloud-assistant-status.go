package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errHandler"
)

// 获取指定实例中的云助手状态，true 表示正常工作
func DescribeCloudAssistantStatus(instanceId string) (bool, *errHandler.CustomErr) {
	client, customErr := aliyun.CreateEcsClient()

	if customErr != nil {
		return false, customErr
	}

	res, err := client.DescribeCloudAssistantStatus(&ecs.DescribeCloudAssistantStatusRequest{
		RegionId:   tea.String(aliyun.AliyunConfig.PrimaryRegionId),
		InstanceId: []*string{tea.String(instanceId)},
	})

	if err != nil {
		return false, errHandler.AliyunError(err)
	}

	for _, item := range res.Body.InstanceCloudAssistantStatusSet.InstanceCloudAssistantStatus {
		return tea.StringValue(item.CloudAssistantStatus) == "true", nil
	}

	return false, errHandler.ResNotExist()
}
