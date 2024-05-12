package coins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/db"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

//go:generate mockgen -package=coins -destination=../../mocks/usecases/coins/init.go -source=init.go
type (
	Usecases interface {
		SyncPrice(ctx context.Context, req dto.SyncCoinPriceRequest)
	}

	usecases struct {
		persistents persistents.Persistents
		log         logger.Logger
		orm         *db.DB
	}
)

func Init(persist persistents.Persistents, log logger.Logger, orm *db.DB) (Usecases, error) {
	return &usecases{
		persistents: persist,
		log:         log,
		orm:         orm,
	}, nil
}
