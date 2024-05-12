package userCoins

import (
	"testing"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_Create(t *testing.T) {
	var userCoin = dao.UserCoin{}

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Create(&userCoin).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.Create(mock.ctx, userCoin)

		assert.Nil(t, err)
		assert.Equal(t, userCoin, res)
	})

	t.Run("If error exist, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Create(&userCoin).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)

		mock.log.EXPECT().Error(mock.ctx, "create-user-coin").Return(mock.log)
		mock.log.EXPECT().Err(errAny).Return(mock.log)
		mock.log.EXPECT().Any("user-coin", userCoin).Return(mock.log)
		mock.log.EXPECT().Send()

		res, err := mock.repositories.Create(mock.ctx, userCoin)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.UserCoin{}, res)
	})
}
