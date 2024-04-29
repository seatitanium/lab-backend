package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errHandler"
)

// 获取指定 invokeId 的执行结果
func DescribeInvocationResults(invokeId string) (string, *errHandler.CustomErr) {
	client, customErr := aliyun.CreateEcsClient()

	if customErr != nil {
		return "", customErr
	}

	res, err := client.DescribeInvocationResults(&ecs.DescribeInvocationResultsRequest{
		RegionId:        tea.String(aliyun.AliyunConfig.PrimaryRegionId),
		InvokeId:        tea.String(invokeId),
		ContentEncoding: tea.String("PlainText"),
	})

	if err != nil {
		return "", errHandler.AliyunError(err)
	}

	for _, item := range res.Body.Invocation.InvocationResults.InvocationResult {
		return tea.StringValue(item.Output), nil
	}

	return "", errHandler.ResNotExist()
}
