package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type EnvConfig struct {
	Name string `envconfig:"NAME" default:"Coincap Cryptoprice API" required:"true"`
	Port int    `envconfig:"PORT" default:"8090" required:"true"`

	DbPath string `envconfig:"DB_PATH" required:"true"`

	DbMaxIdleConnection int           `envconfig:"DB_MAX_IDLE_CONNECTION" default:"10"`
	DbMaxOpenConnection int           `envconfig:"DB_MAX_OPEN_CONNECTION" default:"10"`
	DbConnMaxLifetime   time.Duration `envconfig:"DB_CONNECTION_MAX_LIFE_TIME" default:"60s"`

	DbLogMode          bool          `envconfig:"DB_LOG_MODE" default:"false"`
	DbLogSlowThreshold time.Duration `envconfig:"DB_LOG_SLOW_THRESHOLD"`

	DbRetry     int           `envconfig:"DB_RETRY" default:"3"`
	DbWaitSleep time.Duration `envconfig:"DB_WAIT_SLEEP" default:"1s"`

	LogMode bool `envconfig:"LOG_MODE" default:"false"`

	BcryptCost int `envconfig:"BCRYPT_COST" default:"5" required:"true"`

	JwtPublicKey string        `envconfig:"JWT_PUBLIC_KEY" required:"true"`
	JwtAlg       string        `envconfig:"JWT_ALG" required:"true" default:"HS256"`
	JwtExpires   time.Duration `envconfig:"JWT_EXPIRES" required:"true" default:"1h"`

	CoincapBasePath      string `envconfig:"COINCAP_BASE_PATH" required:"true"`
	CoincapWebsocketPath string `envconfig:"COINCAP_WEBSOCKET_PATH" required:"true"`
	CoincapPriceSyncMode bool   `envconfig:"COINCAP_PRICE_SYNC_MODE" default:"false" required:"true"`
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

		return &config, nil
	}

	if err := godotenv.Load(filename); err != nil {
		return nil, errors.Wrap(err, "failed to read from .env file")
	}

	if err := envconfig.Process("", &config); err != nil {
		return nil, errors.Wrap(err, "failed to read from env variable")
	}

	return &config, nil
}
