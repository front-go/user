package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Postgres struct {
	Host     string `env:"PG_HOST"`
	Port     string `env:"PG_PORT"`
	User     string `env:"PG_USER"`
	Password string `env:"PG_PASSWORD"`
	Database string `env:"PG_DATABASE_NAME"`
}

type Service struct {
	Port string `env:"AUTH_SERVICE_PORT"`
}

type Config struct {
	Postgres Postgres
	Service  Service
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf(".env file not found, using system environment variables")
	}

	cfg := &Config{}
	if err := cleanenv.ReadEnv(cfg); err != nil {
		log.Fatalf("Failed to read env: %v", err)
	}

	return cfg
}
