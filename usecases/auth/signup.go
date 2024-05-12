package auth

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/consts"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/orm/driver/sqliteOrm"
)

func (u usecases) SignUp(ctx context.Context, req dto.SignUpRequest) (dao.User, error) {
	if req.Password != req.PasswordConfirmation {
		return dao.User{}, consts.ErrPasswordConfirmationNotSame
	}

	hashedPass, err := u.pwd.Generate(req.Password)
	if err != nil {
		return dao.User{}, err
	}

	req.Password = hashedPass
	user, err := u.persistents.Repositories.Users.Create(ctx, req.ToUser(ctx))
	if err != nil {
		if sqliteOrm.IsUniqueConstraintError(err) {
			return dao.User{}, consts.ErrEmailHasBeenUsed
		}

		return dao.User{}, err
	}

	return user, nil
}
