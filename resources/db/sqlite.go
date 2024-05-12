package db

import (
	"github.com/naufalfmm/cryptocurrency-price-api/resources/config"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm/driver/sqliteOrm"
)

func NewSqlite(config *config.EnvConfig, log logger.Logger) (*DB, error) {
	confs := []sqliteOrm.SqliteConfig{
		sqliteOrm.WithPath(config.DbPath),
		sqliteOrm.WithMaxIdleConnection(config.DbMaxIdleConnection),
		sqliteOrm.WithMaxOpenConnection(config.DbMaxOpenConnection),
		sqliteOrm.WithConnMaxLifeTime(config.DbConnMaxLifetime),
		sqliteOrm.WithRetry(config.DbRetry, config.DbWaitSleep),
	}

	if config.DbLogMode {
		confs = append(confs, sqliteOrm.WithLog(log, config.DbLogSlowThreshold))
	}

	o, err := sqliteOrm.NewSqlite(confs...)
	if err != nil {
		return nil, err
	}

	return &DB{
		Orm: o,
	}, nil
}
