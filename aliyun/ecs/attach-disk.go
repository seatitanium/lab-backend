package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errors"
)

func AttachDisk(instanceId string, diskId string) *errors.CustomErr {
	client, customErr := aliyun.CreateEcsClient()

	if customErr != nil {
		return customErr
	}

	_, err := client.AttachDisk(&ecs.AttachDiskRequest{
		InstanceId: tea.String(instanceId),
		DiskId:     tea.String(diskId),
	})

	if err != nil {
		return errors.AliyunError(err)
	}

	return nil
}
