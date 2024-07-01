package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/aliyun"
	"seatimc/backend/errors"
)

// 为实例分配公网 IP。
//
// 如果分配成功，返回分配的 IP；分配失败返回空字符串。
func AllocatePublicIpAddress(instanceId string) (string, *errors.CustomErr) {
	client, customErr := aliyun.CreateEcsClient()
	if customErr != nil {

		return "", customErr
	}

	resp, err := client.AllocatePublicIpAddress(&ecs.AllocatePublicIpAddressRequest{InstanceId: tea.String(instanceId)})

	if err != nil {
		return "", errors.AliyunError(err)
	}

	return tea.StringValue(resp.Body.IpAddress), nil
}
