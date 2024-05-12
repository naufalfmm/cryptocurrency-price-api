package db

import (
	"github.com/naufalfmm/cryptocurrency-price-api/migrations/resources/config"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm/driver/sqliteOrm"
)

func NewSqlite(config *config.EnvConfig, log logger.Logger) (orm.Orm, error) {
	confs := []sqliteOrm.SqliteConfig{
		sqliteOrm.WithPath(config.DbPath),
		sqliteOrm.WithRetry(config.DbRetry, config.DbWaitSleep),
	}

	o, err := sqliteOrm.NewSqlite(confs...)
	if err != nil {
		return nil, err
	}

	return o, nil
}
