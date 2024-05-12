package auth

import "github.com/naufalfmm/cryptocurrency-price-api/usecases"

type Controllers struct {
	Usecases usecases.Usecases
}

func Init(usec usecases.Usecases) (Controllers, error) {
	return Controllers{
		Usecases: usec,
	}, nil
}
