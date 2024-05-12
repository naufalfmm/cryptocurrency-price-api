package auth

import (
	"context"

	"github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	"github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/password"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/token/jwt"
)

//go:generate mockgen -package=auth -destination=../../mocks/usecases/auth/init.go -source=init.go
type (
	Usecases interface {
		SignUp(ctx context.Context, req dto.SignUpRequest) (dao.User, error)
		SignIn(ctx context.Context, req dto.SignInRequest) (dao.UserSignIn, error)
	}

	usecases struct {
		persistents persistents.Persistents
		log         logger.Logger
		pwd         password.Password
		jwt         jwt.JWT
	}
)

func Init(persist persistents.Persistents, log logger.Logger, pwd password.Password, jwt jwt.JWT) (Usecases, error) {
	return &usecases{
		persistents: persist,
		log:         log,
		pwd:         pwd,
		jwt:         jwt,
	}, nil
}
