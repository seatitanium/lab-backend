package ecs

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

func CreateClient() (*ecs.Client, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(AConf().AccessKeyId),
		AccessKeySecret: tea.String(AConf().AccessKeySecret),
		RegionId:        tea.String(AConf().PrimaryRegionId),
	}

	config.Endpoint = tea.String("ecs." + AConf().PrimaryRegionId + ".aliyuncs.com")
	return ecs.NewClient(config)
}
