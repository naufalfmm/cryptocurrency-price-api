package usecases

import (
	"github.com/naufalfmm/cryptocurrency-price-api/persistents"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/config"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/db"
	"github.com/naufalfmm/cryptocurrency-price-api/usecases/auth"
	"github.com/naufalfmm/cryptocurrency-price-api/usecases/coins"
	"github.com/naufalfmm/cryptocurrency-price-api/usecases/userCoins"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/password"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/token/jwt"
)

type Usecases struct {
	Auth      auth.Usecases
	Coins     coins.Usecases
	UserCoins userCoins.Usecases
}

func Init(persist persistents.Persistents, p password.Password, j jwt.JWT, l logger.Logger, orm *db.DB, conf *config.EnvConfig) (Usecases, error) {
	a, err := auth.Init(persist, l, p, j)
	if err != nil {
		return Usecases{}, err
	}

	c, err := coins.Init(persist, l, orm)
	if err != nil {
		return Usecases{}, err
	}

	uc, err := userCoins.Init(persist, conf, l, orm)
	if err != nil {
		return Usecases{}, err
	}

	return Usecases{
		Auth:      a,
		Coins:     c,
		UserCoins: uc,
	}, nil
}
