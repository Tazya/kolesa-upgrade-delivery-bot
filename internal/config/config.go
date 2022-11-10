package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Addr     string
	Env      string
	BotToken string
	Dsn      string
}

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
