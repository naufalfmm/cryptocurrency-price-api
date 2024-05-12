package userCoins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/db"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm/driver/sqliteOrm"
)

//go:generate mockgen -package=userCoins -destination=../../../mocks/persistents/repositories/userCoins/init.go -source=init.go
type (
	Repositories interface {
		Create(ctx context.Context, userCoin dao.UserCoin) (dao.UserCoin, error)
		Get(ctx context.Context, userID, coinID uint64) (dao.UserCoin, error)
		DeleteByID(ctx context.Context, id uint64, deletedBy string) error
		GetAll(ctx context.Context, req dto.GetAllRequest, queryModifier sqliteOrm.QueryModifier) ([]dao.UserCoin, error)
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
