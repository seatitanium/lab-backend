package utils

import (
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
)

type ConfigDatabase struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type ConfigToken struct {
	Expiration int    `yaml:"expiration"`
	PrivateKey string `yaml:"private-key"`
}

type Config struct {
	Domain                 string         `yaml:"domain"`
	AllowedOrigins         []string       `yaml:"allowed-origins"`
	Database               ConfigDatabase `yaml:"database"`
	ServerDatabase         ConfigDatabase `yaml:"server-database"`
	Token                  ConfigToken    `yaml:"token"`
	BindPort               int            `yaml:"bind-port"`
	Version                string         `yaml:"version"`
	EnableConfigWhitelist  bool           `yaml:"enable-config-whitelist"`
	NeedAuthorizeEndpoints []string       `yaml:"need-authorize-endpoints"`
	ServerOnlyEndpoints    []string       `yaml:"server-only-endpoints"`
	AdminOnlyEndpoints     []string       `yaml:"admin-only-endpoints"`
	ServerSecret           string         `yaml:"server-secret"`
	ActiveTerm             int            `yaml:"active-term"`
	Administrators         []string       `yaml:"administrators"`
}

var GlobalConfig *Config

func LoadGlobalConfig(path string) {
	GlobalConfig = &Config{}
	cfgFile, err := os.ReadFile(path)
	MustPanic(err)
	err = yaml.Unmarshal(cfgFile, GlobalConfig)
	MustPanic(err)
	return
}

func GetActiveTerm() *Term {
	for _, term := range GetTerms() {
		if term.Tag == "st"+strconv.Itoa(GlobalConfig.ActiveTerm) {
			return &term
		}
	}
	return nil
}
