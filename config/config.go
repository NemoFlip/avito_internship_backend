package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		HTTP `yaml:"http"`
	}

	HTTP struct {
		Port string `yaml:"port" env:"HTTP_PORT" env-default:"8080"`
	}
)

func NewConfig() (*Config, error) {
	var cfg Config
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	if err := cleanenv.ReadConfig("./config/config.yml", &cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	return &cfg, nil
}
