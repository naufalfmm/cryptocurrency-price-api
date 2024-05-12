package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

func create() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		unixMilli := time.Now().UnixMilli()

		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		sqlLocation := path.Join("sql")
		if !strings.Contains(wd, "migrations") {
			sqlLocation = path.Join("migrations", sqlLocation)
		}

		sqlLocation = path.Join(wd, sqlLocation)

		if _, err := os.Stat(sqlLocation); os.IsNotExist(err) {
			if err := os.Mkdir(sqlLocation, 0755); err != nil {
				return err
			}
		}

		name := strings.ToLower(ctx.String("name"))
		name = strings.ReplaceAll(name, " ", "_")

		if _, err := os.Create(path.Join(sqlLocation, fmt.Sprintf("%d_%s_migrate.sql", unixMilli, name))); err != nil {
			return err
		}

		if _, err := os.Create(path.Join(sqlLocation, fmt.Sprintf("%d_%s_rollback.sql", unixMilli, name))); err != nil {
			if err := os.Remove(path.Join(sqlLocation, fmt.Sprintf("%d_%s_migrate.sql", unixMilli, name))); err != nil {
				return err
			}
		}

		return nil
	}
}

func Create() *cli.Command {
	return &cli.Command{
		Name:    "create",
		Usage:   "create --name <name>",
		Aliases: []string{"c"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Required: true,
			},
		},
		Action: create(),
	}
}
