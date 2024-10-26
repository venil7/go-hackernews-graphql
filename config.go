package main

import (
	"log"
	"os"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Schema  string `env:"SCHEMA" envDefault:"./schema.graphql"`
	Address string `env:"ADDRESS" envDefault:"localhost:8080"`
	// Database string `env:"DB" envDefault:"database.db"`
}

func (config *Config) GetSchema() (string, error) {
	bytes, err := os.ReadFile(config.Schema)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return string(bytes), nil
}

func GetConfig() (Config, error) {
	return env.ParseAs[Config]()
}
