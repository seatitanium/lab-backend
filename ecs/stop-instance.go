package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

func StopInstance(instanceId string, force bool) error {
	client, err := CreateClient()

	if err != nil {
		return err
	}

	_, err = client.StopInstance(&ecs.StopInstanceRequest{
		InstanceId: tea.String(instanceId),
		ForceStop:  tea.Bool(force),
	})

	if err != nil {
		return err
	}

	return nil
}
