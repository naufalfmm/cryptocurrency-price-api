package userCoins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm/driver/sqliteOrm"
)

func (r repositories) GetAll(ctx context.Context, req dto.GetAllRequest, queryModifier sqliteOrm.QueryModifier) ([]dao.UserCoin, error) {
	var data []dao.UserCoin
	o := r.orm.GetDB().WithContext(ctx)

	if queryModifier != nil {
		o = queryModifier(o)
	}

	if req.UserID != 0 {
		o = o.Where("user_id = ?", req.UserID)
	}

	if err := o.Find(&data).Error(); err != nil {
		r.log.Error(ctx, "get-all-user-coins").Err(err).Any("req", req).Send()
		return nil, err
	}

	return data, nil
}
