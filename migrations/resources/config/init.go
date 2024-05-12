package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type EnvConfig struct {
	DbPath string `envconfig:"DB_PATH" required:"true"`

	DbRetry     int           `envconfig:"DB_RETRY" default:"3"`
	DbWaitSleep time.Duration `envconfig:"DB_WAIT_SLEEP" default:"1s"`
}

func NewConfig() (*EnvConfig, error) {
	var config EnvConfig

	filename := os.Getenv("CONFIG_FILE")

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", &config); err != nil {
			return nil, errors.Wrap(err, "failed to read from env variable")
		}
	}

	if err := godotenv.Load(filename); err != nil {
		return nil, errors.Wrap(err, "failed to read from .env file")
	}

	if err := envconfig.Process("", &config); err != nil {
		return nil, errors.Wrap(err, "failed to read from env variable")
	}

	return &config, nil
}
