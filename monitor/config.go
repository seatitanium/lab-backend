package monitor

import (
	"gopkg.in/yaml.v3"
	"os"
	"seatimc/backend/utils"
)

type MonitorConf struct {
	Enabled bool `yaml:"enabled"`
}

var MonitorConfig *MonitorConf

func LoadMonitorConfig(path string) {
	MonitorConfig = &MonitorConf{}
	cfgFile, err := os.ReadFile(path)
	utils.MustPanic(err)
	ymlErr := yaml.Unmarshal(cfgFile, MonitorConfig)
	utils.MustPanic(ymlErr)
	return
}
