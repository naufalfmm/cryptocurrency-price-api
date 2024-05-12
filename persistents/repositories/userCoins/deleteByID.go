package userCoins

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/frozenTime"
)

func (r repositories) DeleteByID(ctx context.Context, id uint64, deletedBy string) error {
	now := frozenTime.Now(ctx)
	if err := r.orm.GetDB().
		WithContext(ctx).
		Model(&dao.UserCoin{}).
		Where("id = ?", id).
		UpdateColumns(map[string]interface{}{
			"deleted_at":   now,
			"deleted_by":   deletedBy,
			"deleted_unix": now.Unix(),
		}).
		Error(); err != nil {
		r.log.Error(ctx, "delete-user-coin-by-id").Err(err).Uint64("id", id).Str("deleted-by", deletedBy).Send()
		return err
	}

	return nil
}
