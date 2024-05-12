package userCoins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
)

func (r repositories) Get(ctx context.Context, userID, coinID uint64) (dao.UserCoin, error) {
	var uc dao.UserCoin
	if err := r.orm.GetDB().
		WithContext(ctx).
		Where("user_id = ?", userID).
		Where("coin_id = ?", coinID).
		First(&uc).
		Error(); err != nil {
		r.log.Error(ctx, "get-user-coin").Err(err).Uint64("user-id", userID).Uint64("coin-id", coinID).Send()
		return dao.UserCoin{}, err
	}

	return uc, nil
}
