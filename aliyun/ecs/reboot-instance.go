package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errors"
)

func RebootInstance(instanceId string, force bool) *errors.CustomErr {
	client, customErr := aliyun.CreateEcsClient()
	if customErr != nil {
		return customErr
	}

	_, err := client.RebootInstance(&ecs.RebootInstanceRequest{
		InstanceId: tea.String(instanceId),
		ForceStop:  tea.Bool(force),
	})
	if err != nil {
		return errors.AliyunError(err)
	}

	return nil
}
