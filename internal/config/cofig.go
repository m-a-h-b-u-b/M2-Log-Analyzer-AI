//! M2 Log Analyzer AI
//! --------------------------------
//! License : Dual License
//!           - Apache 2.0 for open-source / personal use
//!           - Commercial license required for closed-source use
//! Author  : Md Mahbubur Rahman
//! URL     : https://m-a-h-b-u-b.github.io
//! GitHub  : https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI
//!
//! Module Description:
//! Loads YAML configuration into structured Go types with default handling.

package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerPort int `yaml:"server_port"`
	Workers    int `yaml:"workers"`
	QueueSize  int `yaml:"queue_size"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfg := &Config{}
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
