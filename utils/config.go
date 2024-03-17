package utils

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ConfigDatabase struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
}

type ConfigToken struct {
	Expiration int    `yaml:"expiration"`
	PrivateKey string `yaml:"private-key"`
}

type Config struct {
	Domain                string         `yaml:"domain"`
	AllowedOrigins        []string       `yaml:"allowed-origins"`
	Database              ConfigDatabase `yaml:"database"`
	Token                 ConfigToken    `yaml:"token"`
	BindPort              int            `yaml:"bind-port"`
	Version               string         `yaml:"version"`
	EnableConfigWhitelist bool           `yaml:"enable-config-whitelist"`
	NeedAuthorizeHandlers []string       `yaml:"need-authorize-handlers"`
}

// 从 config.yml 中获取数据
func Conf() Config {
	cfg := Config{}
	cfgFile, err := os.ReadFile("./config.yml")
	MustPanic(err)
	ymlErr := yaml.Unmarshal([]byte(cfgFile), &cfg)
	MustPanic(ymlErr)
	return cfg
}
