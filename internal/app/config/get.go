package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func GetConfig(configPath string) (GlobalConfig, error) {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return GlobalConfig{}, fmt.Errorf("read %s : %w", configPath, err)
	}

	var cfg GlobalConfig
	if unmarshalErr := yaml.Unmarshal(file, &cfg); unmarshalErr != nil {
		return GlobalConfig{}, fmt.Errorf("unmarshal config: %w", err)
	}

	return cfg, nil
}
