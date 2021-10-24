package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	Env           string
	Debug         bool   `default:"true"`
	Server        ServerConfig
}

func loadConfig() (Config, error) {
	config := Config{}
	_ = godotenv.Load() // nolint
	err := envconfig.Process("", &config)

	return config, errors.Wrap(err, "failed config load")
}
