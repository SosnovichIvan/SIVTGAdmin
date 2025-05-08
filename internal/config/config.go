package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Configuration struct {
		App APP `yaml:"app"`
		Bot BOT `yaml:"bot"`
	}

	APP struct {
		DebugMode bool `yaml:"debug_mode" default:"false" validate:"required"`
	}

	BOT struct {
		Token string `yaml:"token" validate:"required"`
	}
)

const fileName = "config.yaml"

var config *Configuration

func InitConfig() (*Configuration, error) {
	if config != nil {
		return config, nil
	}

	return parseConfig()

}

func GetConfig() *Configuration {
	if config == nil {
		panic("Config not found")
	}

	return config
}

func parseConfig() (*Configuration, error) {
	fmt.Println("Start parse config")

	yamlFile, err := os.ReadFile(fileName)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ошибка чтения файла config.yaml: %v", err))
	}

	cfg := new(Configuration)

	// Парсим YAML в структуру
	err = yaml.Unmarshal(yamlFile, &cfg)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ошибка парсинга YAML: %v", err))
	}

	config = cfg

	fmt.Println("parse config complete")

	return cfg, nil
}
