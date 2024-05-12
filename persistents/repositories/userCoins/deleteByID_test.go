package userCoins

import (
	"testing"
	"time"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/frozenTime"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_DeleteByID(t *testing.T) {
	var (
		id        uint64 = 1
		deletedBy        = "test"

		now = time.Now()
	)

	t.Run("If no error, it will return nil", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.ctx = frozenTime.Freeze(t, mock.ctx, now)

		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Model(&dao.UserCoin{}).Return(mock.orm)
		mock.orm.EXPECT().Where("id = ?", id).Return(mock.orm)
		mock.orm.EXPECT().UpdateColumns(map[string]interface{}{
			"deleted_at":   now,
			"deleted_by":   deletedBy,
			"deleted_unix": now.Unix(),
		}).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(nil)

		err := mock.repositories.DeleteByID(mock.ctx, id, deletedBy)

		assert.Nil(t, err)
	})

	t.Run("If error exist, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.ctx = frozenTime.Freeze(t, mock.ctx, now)

		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Model(&dao.UserCoin{}).Return(mock.orm)
		mock.orm.EXPECT().Where("id = ?", id).Return(mock.orm)
		mock.orm.EXPECT().UpdateColumns(map[string]interface{}{
			"deleted_at":   now,
			"deleted_by":   deletedBy,
			"deleted_unix": now.Unix(),
		}).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)

		mock.log.EXPECT().Error(mock.ctx, "delete-user-coin-by-id").Return(mock.log)
		mock.log.EXPECT().Err(errAny).Return(mock.log)
		mock.log.EXPECT().Uint64("id", id).Return(mock.log)
		mock.log.EXPECT().Str("deleted-by", deletedBy).Return(mock.log)
		mock.log.EXPECT().Send()

		err := mock.repositories.DeleteByID(mock.ctx, id, deletedBy)

		assert.Equal(t, errAny, err)
	})
}
