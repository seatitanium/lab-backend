package aliyun

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type InstanceDiskConf struct {
	// 硬盘种类
	Category string `yaml:"category"`
	// 系统盘名
	DiskName string `yaml:"disk-name"`
	// 硬盘大小
	Size int32 `yaml:"size"`
}

type AliyunUsingConf struct {
	// 实例类型，决定配置。如 ecs.g6.large
	InstanceType string `yaml:"instance-type"`
	// 网络类型
	NetworkType string `yaml:"network-type"`
	// 是否为 IO 优化实例
	IoOptimized bool `yaml:"io-optimized"`
	// 竞价时长。无特殊情况设置为 1
	SpotDuration int32 `yaml:"spot-duration"`
	// 系统类型，windows 或者 linux
	OSType string `yaml:"os-type"`
	// 使用的镜像名称
	ImageId string `yaml:"image-id"`
	// 安全组 ID
	SecurityGroupId string `yaml:"security-group-id"`
	// 实例名称
	InstanceName string `yaml:"instance-name"`
	// 网络计费类型
	InternetChargeType string `yaml:"internet-charge-type"`
	// 公网出带宽最大值
	InternetMaxBandwidthOut int32 `yaml:"internet-max-bandwidth-out"`
	// 实例密码
	Password string `yaml:"password"`
	// 硬盘相关配置
	Disk InstanceDiskConf `yaml:"disk"`
	// 实例付费方式
	InstanceChargeType string `yaml:"instance-charge-type"`
	// 实例抢占策略
	SpotStrategy string `yaml:"spot-strategy"`
	// 是否为预检
	DryRun bool `yaml:"dry-run"`
}

type AliyunConf struct {
	// AKID
	AccessKeyId string `yaml:"access-key-id"`
	// AKSecret
	AccessKeySecret string `yaml:"access-key-secret"`
	// 首选地域 ID，如 cn-shenzhen
	PrimaryRegionId string `yaml:"primary-region-id"`
	// 用于部署的云助手指令 ID，格式 c-xxxxx
	DeployCommandId string `yaml:"deploy-command-id"`
	// 当前使用的实例相关配置
	Using AliyunUsingConf `yaml:"using"`
}

var AliyunConfig *AliyunConf

func LoadAliyunConfig(path string) {
	AliyunConfig = &AliyunConf{}
	cfgFile, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err.Error())
	}

	ymlErr := yaml.Unmarshal(cfgFile, AliyunConfig)

	if ymlErr != nil {
		log.Fatal(ymlErr.Error())
	}
	return
}
