package cmd

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/migrations/model"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm"
)

func checkCreateMigrationLog(ctx context.Context, o orm.Orm) error {
	migrator := o.Migrator()

	if migrator.HasTable(model.MigrationLog{}) {
		return nil
	}

	return migrator.CreateTable(model.MigrationLog{})
}

func checkConnection(ctx context.Context, o orm.Orm) error {
	if err := o.Ping(); err != nil {
		return err
	}

	if err := checkCreateMigrationLog(ctx, o); err != nil {
		return err
	}

	return o.Ping()
}
