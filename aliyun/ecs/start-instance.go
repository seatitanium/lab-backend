package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errHandler"
)

func StartInstance(instanceId string) *errHandler.CustomErr {
	client, customErr := aliyun.CreateEcsClient()

	if customErr != nil {
		return customErr
	}

	_, err := client.StartInstance(&ecs.StartInstanceRequest{
		InstanceId: tea.String(instanceId),
	})

	if err != nil {
		return errHandler.AliyunError(err)
	}

	return nil
}
