package config

import (
	"os"

	"github.com/vanvanni/sudjoo/internal/logger"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Token string `yaml:"token"`
}

var instance *Config

func GetConfig() *Config {
	if instance == nil {
		data, err := os.ReadFile("config.yml")
		if err != nil {
			logger.Fatal("Config: " + err.Error())
		}

		if err := yaml.Unmarshal(data, &instance); err != nil {
			logger.Fatal("Config: " + err.Error())
		}
	} else {
		logger.Debug("Fetched config.go, using memory instance")
	}

	return instance
}
