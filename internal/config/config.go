package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port string `envconfig:"PORT"`
}

func New() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}
	return &c, nil
}
