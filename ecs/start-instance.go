package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

func StartInstance(instanceId string) error {
	client, err := CreateClient()

	if err != nil {
		return err
	}

	_, err = client.StartInstance(&ecs.StartInstanceRequest{
		InstanceId: tea.String(instanceId),
	})

	if err != nil {
		return err
	}

	return nil
}
