package coinHistories

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
)

func (r repositories) BulkCreate(ctx context.Context, coinHistories []dao.CoinHistory) ([]dao.CoinHistory, error) {
	if err := r.orm.GetDB().
		WithContext(ctx).
		Create(&coinHistories).
		Error(); err != nil {
		r.log.Error(ctx, "create-coin-histories").Err(err).Any("coin-histories", coinHistories).Send()
		return nil, err
	}

	return coinHistories, nil
}
