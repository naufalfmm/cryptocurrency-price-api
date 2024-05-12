package controllers

import (
	"github.com/naufalfmm/cryptocurrency-price-api/infrastructures/rest/controllers/auth"
	"github.com/naufalfmm/cryptocurrency-price-api/infrastructures/rest/controllers/userCoins"
	"github.com/naufalfmm/cryptocurrency-price-api/usecases"
)

type Controllers struct {
	Auth      auth.Controllers
	UserCoins userCoins.Controllers
}

func Init(usec usecases.Usecases) (Controllers, error) {
	a, err := auth.Init(usec)
	if err != nil {
		return Controllers{}, err
	}

	uc, err := userCoins.Init(usec)
	if err != nil {
		return Controllers{}, err
	}

	return Controllers{
		Auth:      a,
		UserCoins: uc,
	}, nil
}
