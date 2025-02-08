package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"time"
)

type Config struct {
	Env           string        `env:"ENV"`
	Port          string        `env:"PORT"`
	Timeout       time.Duration `env:"TIMEOUT"`
	IdleTimeout   time.Duration `env:"IDLE_TIMEOUT"`
	ChromeAddress string        `env:"CHROME_ADDRESS"`
}

func MustLoad() *Config {

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Cannot read env: %s", err)
	}

	return &cfg
}
