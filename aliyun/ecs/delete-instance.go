package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errors"
)

func DeleteInstance(instanceId string, force bool) *errors.CustomErr {
	client, customErr := aliyun.CreateEcsClient()
	if customErr != nil {
		return customErr
	}

	_, err := client.DeleteInstance(&ecs.DeleteInstanceRequest{
		InstanceId: tea.String(instanceId),
		Force:      tea.Bool(force),
	})

	if err != nil {
		return errors.AliyunError(err)
	}

	return nil
}
