package coins

import (
	"github.com/naufalfmm/cryptocurrency-price-api/usecases"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

type Listeners struct {
	Usecases usecases.Usecases

	log logger.Logger
}

func Init(usec usecases.Usecases, log logger.Logger) (Listeners, error) {
	return Listeners{
		Usecases: usec,
		log:      log,
	}, nil
}
