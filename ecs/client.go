package ecs

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

func CreateClient() (*ecs.Client, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(Conf().AccessKeyId),
		AccessKeySecret: tea.String(Conf().AccessKeySecret),
		RegionId:        tea.String(Conf().PrimaryRegionId),
	}

	config.Endpoint = tea.String("ecs." + Conf().PrimaryRegionId + ".aliyuncs.com")
	return ecs.NewClient(config)
}
