package sqliteOrm

import (
	"time"

	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger/zeroLogger"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type sqliteConfig struct {
	path string

	maxIdleConnection int
	maxOpenConnection int
	connMaxLifetime   time.Duration

	logger           logger.Logger
	logMode          bool
	logSlowThreshold time.Duration

	retry     int
	waitSleep time.Duration
}

func generateDefault() (sqliteConfig, error) {
	log, err := zeroLogger.NewZeroLogger()
	if err != nil {
		return sqliteConfig{}, err
	}

	return sqliteConfig{
		maxIdleConnection: 10,
		maxOpenConnection: 200,
		connMaxLifetime:   time.Hour,

		logger:           log,
		logMode:          false,
		logSlowThreshold: 200 * time.Millisecond,
	}, nil
}

func (c sqliteConfig) ToGormConfig() *gorm.Config {
	conf := gorm.Config{}

	if c.logMode {
		logConf := gormLogger.Config{
			Colorful: true,
			LogLevel: gormLogger.Info,
		}

		if c.logSlowThreshold != 0 {
			logConf.SlowThreshold = c.logSlowThreshold
		}

		conf.Logger = gormLogger.New(c.logger, logConf)
	}

	return &conf
}

type SqliteConfig func(c *sqliteConfig)

func WithPath(path string) SqliteConfig {
	return func(c *sqliteConfig) {
		c.path = path
	}
}

func WithMaxIdleConnection(maxIdleConnection int) SqliteConfig {
	return func(c *sqliteConfig) {
		c.maxIdleConnection = maxIdleConnection
	}
}

func WithMaxOpenConnection(maxOpenConnection int) SqliteConfig {
	return func(c *sqliteConfig) {
		c.maxOpenConnection = maxOpenConnection
	}
}

func WithConnMaxLifeTime(connMaxLifeTime time.Duration) SqliteConfig {
	return func(c *sqliteConfig) {
		c.connMaxLifetime = connMaxLifeTime
	}
}

func WithConnectionOptions(maxIdleConnection, maxOpenConnection int, maxLifeTime time.Duration) SqliteConfig {
	return func(c *sqliteConfig) {
		c.maxIdleConnection = maxIdleConnection
		c.maxOpenConnection = maxOpenConnection
		c.connMaxLifetime = maxLifeTime
	}
}

func WithLogger(log logger.Logger) SqliteConfig {
	return func(c *sqliteConfig) {
		c.logger = log
		c.logMode = true
	}
}

func WithLog(log logger.Logger, slowThreshold time.Duration) SqliteConfig {
	return func(c *sqliteConfig) {
		c.logger = log
		c.logMode = true
		c.logSlowThreshold = slowThreshold
	}
}

func WithRetry(retry int, waitSleep time.Duration) SqliteConfig {
	return func(c *sqliteConfig) {
		c.retry = retry
		c.waitSleep = waitSleep
	}
}
