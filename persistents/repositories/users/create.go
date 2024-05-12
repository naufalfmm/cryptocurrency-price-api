package users

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
)

func (r repositories) Create(ctx context.Context, user dao.User) (dao.User, error) {
	if err := r.orm.
		GetDB().
		WithContext(ctx).
		Create(&user).
		Error(); err != nil {
		r.log.Error(ctx, "create-user").Err(err).Send()
		return dao.User{}, err
	}

	return user, nil
}
