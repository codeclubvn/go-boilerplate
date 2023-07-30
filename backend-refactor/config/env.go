package config

import (
	"fmt"
	"github.com/caarlos0/env"
)

type Config struct {
	DBHost     string `env:"DB_HOST" envDefault:"DB_HOST"`
	DBPort     string `env:"DB_PORT" envDefault:"DB_PORT"`
	DBUser     string `env:"DB_USER" envDefault:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME" envDefault:"DB_NAME"`
}

var cfg Config

func SetEnv() {
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("Failed to read environment variables: %v", err)
		return
	}
}

func GetEnv() Config {
	return cfg
}
