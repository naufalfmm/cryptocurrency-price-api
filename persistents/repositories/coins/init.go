package coins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/db"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

//go:generate mockgen -package=coins -destination=../../../mocks/persistents/repositories/coins/init.go -source=init.go
type (
	Repositories interface {
		GetByCode(ctx context.Context, code string) (dao.Coin, error)
		Create(ctx context.Context, coin dao.Coin) (dao.Coin, error)
		GetByCoincapIDs(ctx context.Context, coincapIDs []string) ([]dao.Coin, error)
		UpdatePrices(ctx context.Context, coins []dao.Coin) error
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
