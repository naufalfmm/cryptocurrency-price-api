package userCoins

import (
	"testing"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_Get(t *testing.T) {
	var (
		userID uint64 = 1
		coinID uint64 = 1

		userCoin = dao.UserCoin{
			UserID: userID,
			CoinID: coinID,
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		var data dao.UserCoin
		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Where("user_id = ?", userID).Return(mock.orm)
		mock.orm.EXPECT().Where("coin_id = ?", coinID).Return(mock.orm)
		mock.orm.EXPECT().First(&data).DoAndReturn(func(c *dao.UserCoin, conds ...interface{}) interface{} {
			*c = userCoin
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.Get(mock.ctx, userID, coinID)

		assert.Nil(t, err)
		assert.Equal(t, userCoin, res)
	})

	t.Run("If error no exist, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		var data dao.UserCoin
		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Where("user_id = ?", userID).Return(mock.orm)
		mock.orm.EXPECT().Where("coin_id = ?", coinID).Return(mock.orm)
		mock.orm.EXPECT().First(&data).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)

		mock.log.EXPECT().Error(mock.ctx, "get-user-coin").Return(mock.log)
		mock.log.EXPECT().Err(errAny).Return(mock.log)
		mock.log.EXPECT().Uint64("user-id", userID).Return(mock.log)
		mock.log.EXPECT().Uint64("coin-id", coinID).Return(mock.log)
		mock.log.EXPECT().Send()

		res, err := mock.repositories.Get(mock.ctx, userID, coinID)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.UserCoin{}, res)
	})
}
