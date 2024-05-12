package coins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
)

func (r repositories) GetByCoincapIDs(ctx context.Context, coincapIDs []string) ([]dao.Coin, error) {
	var coins []dao.Coin

	if err := r.orm.GetDB().
		WithContext(ctx).
		Where("coincap_id IN (?)", coincapIDs).
		Find(&coins).
		Error(); err != nil {
		r.log.Error(ctx, "get-coins-by-coincap-ids").Err(err).Any("coincap-ids", coincapIDs).Send()
		return nil, err
	}

	return coins, nil
}
