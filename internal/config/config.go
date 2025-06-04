package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Services services `yaml:"services"`
}

type kafkaInfo struct {
	Url []string `yaml:"url"`
}

type services struct {
	Auth       serviceInfo    `yaml:"auth"`
	Users      serviceInfo    `yaml:"users"`
	ApiGateway apiGatewayInfo `yaml:"api-gateway"`
	Kafka      kafkaInfo      `yaml:"kafka"`
	Match      serviceInfo    `yaml:"match"`
	Deck       serviceInfo    `yaml:"deck"`
}

type apiGatewayInfo struct {
	Addr string `yaml:"addr"`
}

type serviceInfo struct {
	GRPCPort string       `yaml:"grpcPort"`
	HttpPort string       `yaml:"httpPort"`
	Database databaseInfo `yaml:"database"`
}

type databaseInfo struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func MustLoad() *Config {
	cfgPath := os.Getenv("CONFIG_PATH")

	if cfgPath == "" {
		panic("CONFIG_PATH environment variable not set")
	}

	file, err := os.ReadFile(cfgPath)

	if err != nil {
		panic(err)
	}

	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)

	return &cfg
}
