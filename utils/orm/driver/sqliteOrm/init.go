package sqliteOrm

import (
	"time"

	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqlite(configs ...SqliteConfig) (orm.Orm, error) {
	conf, err := generateDefault()
	if err != nil {
		return nil, err
	}

	for _, config := range configs {
		config(&conf)
	}

	var db *gorm.DB

	gormConf := conf.ToGormConfig()

	for i := 0; i < conf.retry; i++ {
		db, err = gorm.Open(sqlite.Open(conf.path), gormConf)
		if err == nil {
			break
		}

		time.Sleep(conf.waitSleep)
	}

	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(conf.maxIdleConnection)
	sqlDb.SetMaxOpenConns(conf.maxOpenConnection)
	sqlDb.SetConnMaxLifetime(conf.connMaxLifetime)

	o, err := NewOrm(db, orm.Sqlite)
	if err != nil {
		return nil, err
	}

	return &o, nil
}
