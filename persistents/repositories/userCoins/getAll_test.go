package userCoins

import (
	"testing"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_GetAll(t *testing.T) {
	var (
		userID uint64 = 1

		req = dto.GetAllRequest{
			UserID: userID,
		}

		userCoins = []dao.UserCoin{
			{
				UserID: userID,
			},
		}

		queryModifier = func(o orm.Orm) orm.Orm {
			return o.Joins("Any")
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		var data []dao.UserCoin
		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Joins("Any").Return(mock.orm)
		mock.orm.EXPECT().Where("user_id = ?", userID).Return(mock.orm)
		mock.orm.EXPECT().Find(&data).DoAndReturn(func(c *[]dao.UserCoin, conds ...interface{}) interface{} {
			*c = userCoins
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.GetAll(mock.ctx, req, queryModifier)

		assert.Nil(t, err)
		assert.Equal(t, userCoins, res)
	})

	t.Run("If error no exist, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		var data []dao.UserCoin
		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Joins("Any").Return(mock.orm)
		mock.orm.EXPECT().Where("user_id = ?", userID).Return(mock.orm)
		mock.orm.EXPECT().Find(&data).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)

		mock.log.EXPECT().Error(mock.ctx, "get-all-user-coins").Return(mock.log)
		mock.log.EXPECT().Err(errAny).Return(mock.log)
		mock.log.EXPECT().Any("req", req).Return(mock.log)
		mock.log.EXPECT().Send()

		res, err := mock.repositories.GetAll(mock.ctx, req, queryModifier)

		assert.Equal(t, errAny, err)
		assert.Nil(t, res)
	})
}
