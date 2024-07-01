package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errors"
)

type InvocationResult struct {
	Status       string `json:"status"`
	Content      string `json:"content"`
	ErrorInfo    string `json:"errorInfo"`
	StartTime    string `json:"startTime"`
	FinishedTime string `json:"finishedTime"`
}

// 获取指定 invokeId 的执行结果
func DescribeInvocationResults(invokeId string) (*InvocationResult, *errors.CustomErr) {
	client, customErr := aliyun.CreateEcsClient()

	if customErr != nil {
		return nil, customErr
	}

	res, err := client.DescribeInvocationResults(&ecs.DescribeInvocationResultsRequest{
		RegionId:        tea.String(aliyun.AliyunConfig.PrimaryRegionId),
		InvokeId:        tea.String(invokeId),
		ContentEncoding: tea.String("PlainText"),
	})

	if err != nil {
		return nil, errors.AliyunError(err)
	}

	for _, item := range res.Body.Invocation.InvocationResults.InvocationResult {
		return &InvocationResult{
			Status:       tea.StringValue(item.InvocationStatus),
			Content:      tea.StringValue(item.Output),
			ErrorInfo:    tea.StringValue(item.ErrorInfo),
			StartTime:    tea.StringValue(item.StartTime),
			FinishedTime: tea.StringValue(item.FinishedTime),
		}, nil
	}

	return nil, errors.TargetNotExist()
}
