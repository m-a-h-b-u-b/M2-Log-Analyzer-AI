package config

import (
	"os"

	"gopkg.in/yaml.v3"
	"log"
)

// Config is the global configuration struct
type Config struct {
	HTTP struct {
		Port string `yaml:"port"`
	} `yaml:"http"`

	Pipeline struct {
		WorkerCount int `yaml:"worker_count"`
		QueueSize   int `yaml:"queue_size"`
	} `yaml:"pipeline"`

	Detector struct {
		Type       string  `yaml:"type"`
		WindowSize int     `yaml:"window_size"`
		Threshold  float64 `yaml:"threshold"`
	} `yaml:"detector"`

	Alerts struct {
		SlackWebhook string `yaml:"slack_webhook"`
		Email        string `yaml:"email"`
	} `yaml:"alerts"`
}

// Load reads the YAML config file and unmarshals it
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	log.Printf("Loaded config from %s", path)
	return &cfg, nil
}
