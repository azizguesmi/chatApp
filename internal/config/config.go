package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string `env:"PORT" envDefault:"8080"`
	DB   string `env:"DB" envDefault:"db.chat"`
	JWT  string `env:"JWT" envDefault:"secret"`
}

func NewConfig() (*Config, error) {
	// load .env file
	if err := godotenv.Load("../config/.env"); err != nil {
		log.Println("no .env file found, using system env")
	}

	cfg := &Config{}
	err := env.Parse(cfg)
	return cfg, err
}
