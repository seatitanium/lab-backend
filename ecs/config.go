package ecs

import (
	"gopkg.in/yaml.v3"
	"os"
	"seatimc/backend/utils"
)

type AliyunUsingConf struct {
	// 实例类型，决定配置。如 ecs.g6.large
	InstanceType string `yaml:"instance-type"`
	// 网络类型
	NetworkType string `yaml:"network-type"`
	// 是否为 IO 优化实例
	IoOptimized  bool   `yaml:"io-optimized"`
	SpotDuration int32  `yaml:"spot-duration"`
	OSType       string `json:"os_type"`
}

type AliyunConf struct {
	// AKID
	AccessKeyId string `yaml:"access-key-id"`
	// AKSecret
	AccessKeySecret string `yaml:"access-key-secret"`
	// 首选地域 ID，如 cn-shenzhen
	PrimaryRegionId string `yaml:"primary-region-id"`
	// 当前使用的实例相关配置
	Using AliyunUsingConf `yaml:"using"`
}

// 从 aconfig.yml 中获取数据
func AConf() AliyunConf {
	cfg := AliyunConf{}
	cfgFile, err := os.ReadFile("./aconfig.yml")
	utils.MustPanic(err)
	ymlErr := yaml.Unmarshal(cfgFile, &cfg)
	utils.MustPanic(ymlErr)
	return cfg
}
