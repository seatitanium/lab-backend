package lab_backend

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ConfigDatabase struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
}

type ConfigToken struct {
	Expiration int    `yaml:"expiration"`
	PrivateKey string `yaml:"private-key"`
}

type Config struct {
	Domain         string         `yaml:"domain"`
	AllowedOrigins []string       `yaml:"allowed-origins"`
	Database       ConfigDatabase `yaml:"database"`
	Token          ConfigToken    `yaml:"token"`
}

func Conf() Config {
	cfg := Config{}
	cfgFile, err := os.ReadFile("./config.yml")
	MustPanic(err)
	ymlErr := yaml.Unmarshal([]byte(cfgFile), &cfg)
	MustPanic(ymlErr)
	return cfg
}
