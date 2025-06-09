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

type s3Info struct {
	Endpoint  string `yaml:"endpoint"`
	Region    string `yaml:"region"`
	Bucket    string `yaml:"bucket"`
	UseSSl    bool   `yaml:"useSSL"`
	SecretKey string
	AccessKey string
}

type services struct {
	Auth          serviceInfo    `yaml:"auth"`
	Users         serviceInfo    `yaml:"users"`
	ApiGateway    apiGatewayInfo `yaml:"api-gateway"`
	Kafka         kafkaInfo      `yaml:"kafka"`
	Match         serviceInfo    `yaml:"match"`
	Deck          serviceInfo    `yaml:"deck"`
	ObjectStorage s3Info         `yaml:"object-storage"`
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
	s3AccessKey := os.Getenv("S3_ACCESS_KEY_ID")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")

	cfgPath := os.Getenv("CONFIG_PATH")

	if cfgPath == "" || s3AccessKey == "" || s3SecretKey == "" {
		panic("some environment variables are not set: CONFIG_PATH, S3_ACCESS_KEY_ID, S3_SECRET_KEY")
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
	
	cfg.Services.ObjectStorage.SecretKey = s3SecretKey
	cfg.Services.ObjectStorage.AccessKey = s3AccessKey

	fmt.Println(cfg)

	return &cfg
}
