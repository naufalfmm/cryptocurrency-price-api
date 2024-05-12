package coins

import (
	"testing"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_GetByCoincapIDs(t *testing.T) {
	var (
		coincapIDs = []string{
			"ethereum",
		}

		coins = []dao.Coin{
			{
				CoincapID: "ethereum",
			},
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		var data []dao.Coin
		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Where("coincap_id IN (?)", coincapIDs).Return(mock.orm)
		mock.orm.EXPECT().Find(&data).DoAndReturn(func(c *[]dao.Coin, conds ...interface{}) interface{} {
			*c = coins
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.GetByCoincapIDs(mock.ctx, coincapIDs)

		assert.Nil(t, err)
		assert.Equal(t, coins, res)
	})

	t.Run("If error exist, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		var data []dao.Coin
		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Where("coincap_id IN (?)", coincapIDs).Return(mock.orm)
		mock.orm.EXPECT().Find(&data).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)

		mock.log.EXPECT().Error(mock.ctx, "get-coins-by-coincap-ids").Return(mock.log)
		mock.log.EXPECT().Err(errAny).Return(mock.log)
		mock.log.EXPECT().Any("coincap-ids", coincapIDs).Return(mock.log)
		mock.log.EXPECT().Send()

		res, err := mock.repositories.GetByCoincapIDs(mock.ctx, coincapIDs)

		assert.Equal(t, errAny, err)
		assert.Nil(t, res)
	})
}
