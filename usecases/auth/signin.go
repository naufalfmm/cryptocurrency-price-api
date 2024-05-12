package auth

import (
	"context"
	"errors"

	"github.com/naufalfmm/cryptocurrency-price-api/consts"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"gorm.io/gorm"
)

func (u usecases) SignIn(ctx context.Context, req dto.SignInRequest) (dao.UserSignIn, error) {
	user, err := u.persistents.Repositories.Users.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = consts.ErrEmailMissing
		}

		return dao.UserSignIn{}, err
	}

	if err := u.pwd.Check(user.Password, req.Password); err != nil {
		return dao.UserSignIn{}, consts.ErrWrongPassword
	}

	token, err := u.jwt.Encoder.EncodeToken(&dto.LoginClaims{
		UserLogin: dto.UserSignIn{
			ID:    user.ID,
			Email: user.Email,
		},
	})
	if err != nil {
		return dao.UserSignIn{}, err
	}

	return dao.UserSignIn{
		Token: token,
		User: dao.User{
			ID:    user.ID,
			Email: user.Email,
		},
	}, nil
}
