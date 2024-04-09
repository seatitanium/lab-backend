package ecs

import (
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

// 按照 aconfig.yml 中的配置创建一个新的实例，并返回成交价格和实例 ID
func CreateInstance(conf AliyunConf) (*CreatedInstance, error) {
	client, err := CreateClient()

	if err != nil {
		return nil, err
	}

	request := &ecs.CreateInstanceRequest{
		InstanceType:            tea.String(conf.Using.InstanceType),
		IoOptimized:             tea.String(GetIoOptimized(conf.Using.IoOptimized)),
		SpotDuration:            tea.Int32(conf.Using.SpotDuration),
		ImageId:                 tea.String(conf.Using.ImageId),
		SecurityGroupId:         tea.String(conf.Using.SecurityGroupId),
		InstanceName:            tea.String(conf.Using.InstanceName),
		InternetChargeType:      tea.String(conf.Using.InternetChargeType),
		InternetMaxBandwidthOut: tea.Int32(conf.Using.InternetMaxBandwidthOut),
		Password:                tea.String(conf.Using.Password),
		InstanceChargeType:      tea.String(conf.Using.InstanceChargeType),
		SpotStrategy:            tea.String(conf.Using.SpotStrategy),
		DryRun:                  tea.Bool(conf.Using.DryRun),
	}

	request.SystemDisk = &ecs.CreateInstanceRequestSystemDisk{
		DiskName: tea.String(conf.Using.Disk.DiskName),
		Category: tea.String(conf.Using.Disk.Category),
		Size:     tea.Int32(conf.Using.Disk.Size),
	}

	resp, err := client.CreateInstance(request)

	if err != nil {
		return nil, err
	}

	return &CreatedInstance{
		TradePrice: tea.Float32Value(resp.Body.TradePrice),
		InstanceId: tea.StringValue(resp.Body.InstanceId),
	}, nil
}
