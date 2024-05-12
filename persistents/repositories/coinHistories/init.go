package coinHistories

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/db"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

//go:generate mockgen -package=coinHistories -destination=../../../mocks/persistents/repositories/coinHistories/init.go -source=init.go
type (
	Repositories interface {
		BulkCreate(ctx context.Context, coinHistories []dao.CoinHistory) ([]dao.CoinHistory, error)
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
