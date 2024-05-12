package repositories

import (
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/repositories/coinHistories"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/repositories/coins"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/repositories/userCoins"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/repositories/users"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/db"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

type Repositories struct {
	Users         users.Repositories
	Coins         coins.Repositories
	UserCoins     userCoins.Repositories
	CoinHistories coinHistories.Repositories
}

func Init(o *db.DB, l logger.Logger) (Repositories, error) {
	u, err := users.Init(o, l)
	if err != nil {
		return Repositories{}, err
	}

	c, err := coins.Init(o, l)
	if err != nil {
		return Repositories{}, err
	}

	uc, err := userCoins.Init(o, l)
	if err != nil {
		return Repositories{}, err
	}

	ch, err := coinHistories.Init(o, l)
	if err != nil {
		return Repositories{}, err
	}

	return Repositories{
		Users:         u,
		Coins:         c,
		UserCoins:     uc,
		CoinHistories: ch,
	}, nil
}
