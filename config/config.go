package config

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Api Api
}

type Api struct {
	Port    string
	Timeout time.Duration
}

func New(fileName string) (*Config, error) {
	viper.SetConfigFile(fileName)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.WithMessage(err, "read in config")
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.WithMessage(err, "viper unmarshal")
	}

	return cfg, nil
}
