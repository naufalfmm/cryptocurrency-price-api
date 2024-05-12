package userCoins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/config"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/db"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

//go:generate mockgen -package=userCoins -destination=../../mocks/usecases/userCoins/init.go -source=init.go
type (
	Usecases interface {
		TrackCoin(ctx context.Context, req dto.TrackUntrackCoinRequest) (dao.UserCoin, error)
		UntrackCoin(ctx context.Context, req dto.TrackUntrackCoinRequest) error
		GetAllTrack(ctx context.Context, req dto.GetAllTrackRequest) ([]dao.UserCoin, error)
	}

	usecases struct {
		persistents persistents.Persistents
		conf        *config.EnvConfig
		log         logger.Logger
		o           *db.DB
	}
)

func Init(persist persistents.Persistents, conf *config.EnvConfig, log logger.Logger, o *db.DB) (Usecases, error) {
	return &usecases{
		persistents: persist,
		conf:        conf,
		log:         log,
		o:           o,
	}, nil
}
