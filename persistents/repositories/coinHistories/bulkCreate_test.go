package coinHistories

import (
	"testing"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_BulkCreate(t *testing.T) {
	var coinHistories = []dao.CoinHistory{
		{
			CoinID: 1,
		},
		{
			CoinID: 2,
		},
	}

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Create(&coinHistories).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.BulkCreate(mock.ctx, coinHistories)

		assert.Nil(t, err)
		assert.Equal(t, coinHistories, res)
	})

	t.Run("If error exist, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Create(&coinHistories).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)

		mock.log.EXPECT().Error(mock.ctx, "create-coin-histories").Return(mock.log)
		mock.log.EXPECT().Err(errAny).Return(mock.log)
		mock.log.EXPECT().Any("coin-histories", coinHistories).Return(mock.log)
		mock.log.EXPECT().Send()

		res, err := mock.repositories.BulkCreate(mock.ctx, coinHistories)

		assert.Equal(t, errAny, err)
		assert.Nil(t, res)
	})
}
