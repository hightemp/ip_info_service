package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	_ "gopkg.in/yaml.v3"
)

type Config struct {
	Port                 string `yaml:"port"`
	ContriesDataFilePath string `yaml:"countries_data"`
	OrgDataFilePath      string `yaml:"org_data"`
}

var config Config

func Load(f string) error {
	bytes, err := os.ReadFile(f)

	if err != nil {
		return fmt.Errorf("Can't load config: %v", err)
	}

	err = yaml.Unmarshal(bytes, &config)

	if err != nil {
		return fmt.Errorf("Can't parse yaml config: %v", err)
	}

	return nil
}

func CanLoadData() bool {
	return config.ContriesDataFilePath != "" || config.OrgDataFilePath != ""
}

func Get() *Config {
	return &config
}
