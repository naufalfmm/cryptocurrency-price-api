package coins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
)

func (r repositories) Create(ctx context.Context, coin dao.Coin) (dao.Coin, error) {
	if err := r.orm.GetDB().
		WithContext(ctx).
		Create(&coin).
		Error(); err != nil {
		r.log.Error(ctx, "create-coin").Err(err).Any("coin", coin).Send()
		return dao.Coin{}, err
	}

	return coin, nil
}
