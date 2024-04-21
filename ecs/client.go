package ecs

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/errHandler"
)

func CreateClient() (*ecs.Client, *errHandler.CustomErr) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(AliyunConfig.AccessKeyId),
		AccessKeySecret: tea.String(AliyunConfig.AccessKeySecret),
		RegionId:        tea.String(AliyunConfig.PrimaryRegionId),
	}

	config.Endpoint = tea.String("ecs." + AliyunConfig.PrimaryRegionId + ".aliyuncs.com")

	ecsClient, err := ecs.NewClient(config)
	if err != nil {
		return nil, errHandler.AliyunError(err)
	}

	return ecsClient, nil
}
