package users

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
)

func (r repositories) GetByEmail(ctx context.Context, email string) (dao.User, error) {
	var user dao.User
	if err := r.orm.GetDB().
		WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error(); err != nil {
		r.log.Error(ctx, "get-user-by-email").Err(err).Str("email", email).Send()
		return dao.User{}, err
	}

	return user, nil
}
