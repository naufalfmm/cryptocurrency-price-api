package coins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r repositories) UpdatePrices(ctx context.Context, coins []dao.Coin) error {
	if err := r.orm.GetDB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "coincap_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"latest_price": gorm.Expr("excluded.latest_price"),
			"updated_at":   gorm.Expr("excluded.updated_at"),
			"updated_by":   gorm.Expr("excluded.updated_by"),
		}),
	}).Create(&coins).Error(); err != nil {
		r.log.Error(ctx, "update-coin-prices").Err(err).Send()
		return err
	}

	return nil
}
