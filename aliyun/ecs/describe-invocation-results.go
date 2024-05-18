package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errHandler"
	"seatimc/backend/utils"
	"time"
)

type InvocationResult struct {
	Status       string    `json:"status"`
	Content      string    `json:"content"`
	ErrorInfo    string    `json:"errorInfo"`
	StartTime    time.Time `json:"startTime"`
	FinishedTime time.Time `json:"finishedTime"`
}

// 获取指定 invokeId 的执行结果
func DescribeInvocationResults(invokeId string) (*InvocationResult, *errHandler.CustomErr) {
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
		return nil, errHandler.AliyunError(err)
	}

	for _, item := range res.Body.Invocation.InvocationResults.InvocationResult {
		var startTime time.Time
		var finishedTime time.Time

		startTimeStr := tea.StringValue(item.StartTime)
		finishedTimeStr := tea.StringValue(item.FinishedTime)

		startTime, err = utils.ParseTime(startTimeStr)

		if err != nil {
			return nil, errHandler.ServerError(err)
		}

		finishedTime, err = utils.ParseTime(finishedTimeStr)

		if err != nil {
			return nil, errHandler.ServerError(err)
		}

		return &InvocationResult{
			Status:       tea.StringValue(item.InvocationStatus),
			Content:      tea.StringValue(item.Output),
			ErrorInfo:    tea.StringValue(item.ErrorInfo),
			StartTime:    finishedTime,
			FinishedTime: startTime,
		}, nil
	}

	return nil, errHandler.TargetNotExist()
}
