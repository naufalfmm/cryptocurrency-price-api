package coins

import (
	"testing"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_GetByCode(t *testing.T) {
	var (
		code = "ETH"

		coin = dao.Coin{
			Code: code,
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		var data dao.Coin
		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Where("code = ?", code).Return(mock.orm)
		mock.orm.EXPECT().First(&data).DoAndReturn(func(c *dao.Coin, conds ...interface{}) interface{} {
			*c = coin
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.GetByCode(mock.ctx, code)

		assert.Nil(t, err)
		assert.Equal(t, coin, res)
	})

	t.Run("If error no exist, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		var data dao.Coin
		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Where("code = ?", code).Return(mock.orm)
		mock.orm.EXPECT().First(&data).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)

		mock.log.EXPECT().Error(mock.ctx, "get-coin-by-code").Return(mock.log)
		mock.log.EXPECT().Err(errAny).Return(mock.log)
		mock.log.EXPECT().Str("code", code).Return(mock.log)
		mock.log.EXPECT().Send()

		res, err := mock.repositories.GetByCode(mock.ctx, code)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Coin{}, res)
	})
}
