package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Port string `env:"PORT" envDefault:"8080"`
	DB   string `env:"DB" envDefault:"db.chat"`
	JWT  string `env:"JWT" envDefault:"secret"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	return cfg, err
}
