package coincap

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

//go:generate mockgen -package=coincap -destination=../../../mocks/persistents/webclients/coincap/init.go -source=init.go
type (
	Coincap interface {
		GetAllAssets(ctx context.Context, req dto.AllAssetsCoincapRequest) (dao.AllAsset, error)
	}

	coincap struct {
		basePath string
		log      logger.Logger
	}
)

func Init(basePath string, log logger.Logger) (Coincap, error) {
	return &coincap{
		basePath: basePath,
		log:      log,
	}, nil
}
