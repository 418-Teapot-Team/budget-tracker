package config

import (
	"github.com/caarlos0/env"
	"log"
)

type Config struct {
	DatabaseUrl string `env:"DATABASE_URI"`
	Port        string `env:"PORT"`
}

func NewConfig() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}
	return &cfg
}
