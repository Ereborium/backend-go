package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Api Api
}

type Api struct {
	Port    string
	Timeout time.Duration
	Release bool
}

func NewConfig(fileName string) *Config {
	viper.SetConfigFile(fileName)
	viper.SetConfigType("yaml")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &config
}
