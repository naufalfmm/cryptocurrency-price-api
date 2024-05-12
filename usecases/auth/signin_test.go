package auth

import (
	"testing"

	"github.com/naufalfmm/cryptocurrency-price-api/consts"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Test_usecases_SignIn(t *testing.T) {
	var (
		req = dto.SignInRequest{
			Email:    "test@mailinator.com",
			Password: "test123",
		}

		user = dao.User{
			Email:    req.Email,
			Password: "xxxxxxxxxxxxxxx",
		}

		token = "aaaa.bbbb.cccc"
	)

	t.Run("If password match and user exists, it will return sign in data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.user.EXPECT().GetByEmail(mock.ctx, req.Email).Return(user, nil)
		mock.pwd.EXPECT().Check(user.Password, req.Password).Return(nil)
		mock.enc.EXPECT().EncodeToken(&dto.LoginClaims{
			UserLogin: dto.UserSignIn{
				ID:    user.ID,
				Email: user.Email,
			},
		}).Return(token, nil)

		expResp := dao.UserSignIn{
			Token: token,
			User: dao.User{
				ID:    user.ID,
				Email: user.Email,
			},
		}

		resp, err := mock.usecases.SignIn(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, expResp, resp)
	})

	t.Run("If encode token error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.user.EXPECT().GetByEmail(mock.ctx, req.Email).Return(user, nil)
		mock.pwd.EXPECT().Check(user.Password, req.Password).Return(nil)
		mock.enc.EXPECT().EncodeToken(&dto.LoginClaims{
			UserLogin: dto.UserSignIn{
				ID:    user.ID,
				Email: user.Email,
			},
		}).Return("", errAny)

		expResp := dao.UserSignIn{}

		resp, err := mock.usecases.SignIn(mock.ctx, req)

		assert.Equal(t, errAny, err)
		assert.Equal(t, expResp, resp)
	})

	t.Run("If password not match, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.user.EXPECT().GetByEmail(mock.ctx, req.Email).Return(user, nil)
		mock.pwd.EXPECT().Check(user.Password, req.Password).Return(bcrypt.ErrMismatchedHashAndPassword)

		expResp := dao.UserSignIn{}

		resp, err := mock.usecases.SignIn(mock.ctx, req)

		assert.Equal(t, consts.ErrWrongPassword, err)
		assert.Equal(t, expResp, resp)
	})

	t.Run("If get user by email error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.user.EXPECT().GetByEmail(mock.ctx, req.Email).Return(dao.User{}, errAny)

		expResp := dao.UserSignIn{}

		resp, err := mock.usecases.SignIn(mock.ctx, req)

		assert.Equal(t, errAny, err)
		assert.Equal(t, expResp, resp)
	})

	t.Run("If get user by email is not found, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.user.EXPECT().GetByEmail(mock.ctx, req.Email).Return(dao.User{}, gorm.ErrRecordNotFound)

		expResp := dao.UserSignIn{}

		resp, err := mock.usecases.SignIn(mock.ctx, req)

		assert.Equal(t, consts.ErrEmailMissing, err)
		assert.Equal(t, expResp, resp)
	})
}
