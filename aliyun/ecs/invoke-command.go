package ecs

import (
	client2 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errHandler"
	"seatimc/backend/utils"
)

// 在 instanceId 所指向的实例上执行预先设定的云助手指令，并写入 invoke 记录
func InvokeCommand(instanceId string) *errHandler.CustomErr {
	client, customErr := aliyun.CreateEcsClient()

	if customErr != nil {
		return customErr
	}

	res, err := client.InvokeCommand(&client2.InvokeCommandRequest{
		RegionId:   tea.String(aliyun.AliyunConfig.PrimaryRegionId),
		InstanceId: []*string{tea.String(instanceId)},
		CommandId:  tea.String(aliyun.AliyunConfig.DeployCommandId),
	})

	if err != nil {
		return errHandler.AliyunError(err)
	}

	customErr = utils.WriteInvokeRecord(instanceId, tea.StringValue(res.Body.InvokeId))

	if customErr != nil {
		return customErr
	}

	return nil
}
