package users

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/db"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

//go:generate mockgen -package=users -destination=../../../mocks/persistents/repositories/users/init.go -source=init.go
type (
	Repositories interface {
		Create(ctx context.Context, user dao.User) (dao.User, error)
		GetByEmail(ctx context.Context, email string) (dao.User, error)
	}

	repositories struct {
		orm *db.DB
		log logger.Logger
	}
)

func Init(o *db.DB, l logger.Logger) (Repositories, error) {
	return &repositories{
		orm: o,
		log: l,
	}, nil
}
