package users

import (
	"testing"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/stretchr/testify/assert"
)

func Test_repositories_GetByEmail(t *testing.T) {
	var (
		email = "ETH"

		user = dao.User{
			Email: email,
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		var data dao.User
		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Where("email = ?", email).Return(mock.orm)
		mock.orm.EXPECT().First(&data).DoAndReturn(func(c *dao.User, conds ...interface{}) interface{} {
			*c = user
			return mock.orm
		})
		mock.orm.EXPECT().Error().Return(nil)

		res, err := mock.repositories.GetByEmail(mock.ctx, email)

		assert.Nil(t, err)
		assert.Equal(t, user, res)
	})

	t.Run("If error no exist, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		var data dao.User
		mock.orm.EXPECT().WithContext(mock.ctx).Return(mock.orm)
		mock.orm.EXPECT().Where("email = ?", email).Return(mock.orm)
		mock.orm.EXPECT().First(&data).Return(mock.orm)
		mock.orm.EXPECT().Error().Return(errAny)

		mock.log.EXPECT().Error(mock.ctx, "get-user-by-email").Return(mock.log)
		mock.log.EXPECT().Err(errAny).Return(mock.log)
		mock.log.EXPECT().Str("email", email).Return(mock.log)
		mock.log.EXPECT().Send()

		res, err := mock.repositories.GetByEmail(mock.ctx, email)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.User{}, res)
	})
}
