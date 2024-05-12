package userCoins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
)

func (r repositories) Create(ctx context.Context, userCoin dao.UserCoin) (dao.UserCoin, error) {
	if err := r.orm.GetDB().
		WithContext(ctx).
		Create(&userCoin).
		Error(); err != nil {
		r.log.Error(ctx, "create-user-coin").Err(err).Any("user-coin", userCoin).Send()
		return dao.UserCoin{}, err
	}

	return userCoin, nil
}
