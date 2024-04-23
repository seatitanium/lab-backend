package aliyun

import (
	aliyunBss "github.com/alibabacloud-go/bssopenapi-20171214/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiv2 "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	aliyunEcs "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"seatimc/backend/errHandler"
)

func CreateEcsClient() (*aliyunEcs.Client, *errHandler.CustomErr) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(AliyunConfig.AccessKeyId),
		AccessKeySecret: tea.String(AliyunConfig.AccessKeySecret),
		RegionId:        tea.String(AliyunConfig.PrimaryRegionId),
	}

	config.Endpoint = tea.String("ecs." + AliyunConfig.PrimaryRegionId + ".aliyuncs.com")

	ecsClient, err := aliyunEcs.NewClient(config)
	if err != nil {
		return nil, errHandler.AliyunError(err)
	}

	return ecsClient, nil
}

func CreateBssClient() (*aliyunBss.Client, *errHandler.CustomErr) {
	config := &openapiv2.Config{
		AccessKeyId:     tea.String(AliyunConfig.AccessKeyId),
		AccessKeySecret: tea.String(AliyunConfig.AccessKeySecret),
		RegionId:        tea.String(AliyunConfig.PrimaryRegionId),
	}

	config.Endpoint = tea.String("business.aliyuncs.com")

	bssClient, err := aliyunBss.NewClient(config)
	if err != nil {
		return nil, errHandler.AliyunError(err)
	}

	return bssClient, nil
}
