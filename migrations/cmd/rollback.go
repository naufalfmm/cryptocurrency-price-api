package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/naufalfmm/cryptocurrency-price-api/migrations/model"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func isTargetVersionExist(ctx context.Context, o orm.Orm, targetVersion string) (bool, error) {
	if targetVersion == "" {
		return true, nil
	}

	var count int64
	if err := o.WithContext(ctx).
		Model(&model.MigrationLog{}).
		Where("id", targetVersion).
		Count(&count).
		Error(); err != nil {
		return false, err
	}

	return count > 0, nil
}

func rollbackVersion(ctx context.Context, o orm.Orm) (model.MigrationLog, error) {
	var (
		log model.MigrationLog
	)
	if err := o.WithContext(ctx).
		Order("id DESC").
		Take(&log).
		Error(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.MigrationLog{}, nil
		}

		return model.MigrationLog{}, err
	}

	filePath := path.Join(getSQLPath(), fmt.Sprintf("%s_%s_rollback.sql", log.ID, log.Name))
	content, err := os.ReadFile(filePath)
	if err != nil {
		return model.MigrationLog{}, err
	}

	if err := o.WithContext(ctx).
		Exec(string(content)).
		Error(); err != nil {
		return model.MigrationLog{}, err
	}

	if err := o.WithContext(ctx).
		Where("id", log.ID).
		Delete(model.MigrationLog{}).
		Error(); err != nil {
		return model.MigrationLog{}, err
	}

	return log, nil
}

func rollback(o orm.Orm) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		if err := checkConnection(ctx.Context, o); err != nil {
			return err
		}

		ver := ctx.String("version")

		isExist, err := isTargetVersionExist(ctx.Context, o, ver)
		if err != nil {
			return err
		}

		if !isExist {
			return nil
		}

		o.Begin()
		defer o.Rollback()

		log, err := rollbackVersion(ctx.Context, o)
		if err != nil {
			return err
		}
		for log.ID != ver && log.ID != "" {
			log, err = rollbackVersion(ctx.Context, o)
			if err != nil {
				return err
			}
		}

		o.Commit()

		return nil
	}
}

func Rollback(o orm.Orm) *cli.Command {
	return &cli.Command{
		Name:    "rollback",
		Usage:   "rollback --version <version>",
		Aliases: []string{"r"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "version",
				Aliases:  []string{"v"},
				Required: false,
			},
		},
		Action: rollback(o),
	}
}
