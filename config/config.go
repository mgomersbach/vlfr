package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port           string
	MetricsPort    string
	LogPath        string
	LogLevel       string
	Matchers       []string
	MatcherConfig  map[string]map[string]string
	Notifiers      []string
	NotifierConfig map[string]map[string]string
}

func GetConfig(configFile string) (*Config, error) {
	configData, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(configData, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
