package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

func DeleteInstance(instanceId string, force bool) error {
	client, err := CreateClient()

	if err != nil {
		return err
	}

	_, err = client.DeleteInstance(&ecs.DeleteInstanceRequest{
		InstanceId: tea.String(instanceId),
		Force:      tea.Bool(force),
	})

	if err != nil {
		return err
	}

	return nil
}
