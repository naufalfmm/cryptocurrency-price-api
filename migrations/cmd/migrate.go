package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/naufalfmm/cryptocurrency-price-api/migrations/model"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm"
	"github.com/urfave/cli/v2"
)

func generateName(splittedFileNames []string) string {
	var name string
	for _, splittedFileName := range splittedFileNames {
		name = fmt.Sprintf("%s%s_", name, splittedFileName)
	}

	return name[:len(name)-1]
}

func migrate(o orm.Orm) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		if err := checkConnection(ctx.Context, o); err != nil {
			return err
		}

		o.Begin()
		defer o.Rollback()

		err := filepath.Walk(getSQLPath(), func(p string, info fs.FileInfo, _ error) error {
			_, file := filepath.Split(p)
			if file == "" {
				return nil
			}

			splittedFile := strings.Split(strings.TrimSuffix(file, filepath.Ext(file)), "_")

			if splittedFile[len(splittedFile)-1] != "migrate" {
				return nil
			}

			var count int64
			if err := o.WithContext(ctx.Context).
				Model(model.MigrationLog{}).
				Where("id", splittedFile[0]).
				Count(&count).
				Error(); err != nil {
				return err
			}

			if count > 0 {
				return nil
			}

			content, err := os.ReadFile(p)
			if err != nil {
				return err
			}

			if err := o.WithContext(ctx.Context).
				Exec(string(content)).
				Error(); err != nil {
				return err
			}

			if err := o.WithContext(ctx.Context).
				Create(model.MigrationLog{
					ID:        splittedFile[0],
					Name:      generateName(splittedFile[1 : len(splittedFile)-1]),
					MigrateAt: time.Now(),
				}).
				Error(); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}

		o.Commit()

		return nil
	}
}

func Migrate(o orm.Orm) *cli.Command {
	return &cli.Command{
		Name:    "migrate",
		Usage:   "migrate",
		Aliases: []string{"m"},
		Action:  migrate(o),
	}
}
