package coins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
)

func (r repositories) GetByCode(ctx context.Context, code string) (dao.Coin, error) {
	var coin dao.Coin
	if err := r.orm.GetDB().
		WithContext(ctx).
		Where("code = ?", code).
		First(&coin).
		Error(); err != nil {
		r.log.Error(ctx, "get-coin-by-code").Err(err).Str("code", code).Send()
		return dao.Coin{}, err
	}

	return coin, nil
}
