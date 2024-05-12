package listeners

import (
	"github.com/naufalfmm/cryptocurrency-price-api/infrastructures/wslistener/listeners/coins"
	"github.com/naufalfmm/cryptocurrency-price-api/usecases"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

type Listeners struct {
	Coins coins.Listeners
}

func Init(usec usecases.Usecases, log logger.Logger) (Listeners, error) {
	c, err := coins.Init(usec, log)
	if err != nil {
		return Listeners{}, err
	}

	return Listeners{
		Coins: c,
	}, nil
}
