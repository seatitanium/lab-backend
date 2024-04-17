package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/errHandler"
)

func StopInstance(instanceId string, force bool) *errHandler.CustomErr {
	client, customErr := CreateClient()

	if customErr != nil {
		return customErr
	}

	_, err := client.StopInstance(&ecs.StopInstanceRequest{
		InstanceId: tea.String(instanceId),
		ForceStop:  tea.Bool(force),
	})
	if err != nil {
		return errHandler.AliyunError(err)
	}

	return nil
}
