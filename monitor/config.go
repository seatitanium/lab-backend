package monitor

import (
	"gopkg.in/yaml.v3"
	"os"
	"seatimc/backend/utils"
)

type MonitorConf struct {
	Enabled bool `yaml:"enabled"`
}

// 从 aconfig.yml 中获取数据
func Conf() MonitorConf {
	cfg := MonitorConf{}
	cfgFile, err := os.ReadFile("./monitor.yml")
	utils.MustPanic(err)
	ymlErr := yaml.Unmarshal(cfgFile, &cfg)
	utils.MustPanic(ymlErr)
	return cfg
}
