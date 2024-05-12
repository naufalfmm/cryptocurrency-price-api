package persistents

import (
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/repositories"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/webclients"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/config"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/db"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

type Persistents struct {
	Repositories repositories.Repositories
	Webclients   webclients.Webclients
}

func Init(o *db.DB, l logger.Logger, conf *config.EnvConfig) (Persistents, error) {
	repo, err := repositories.Init(o, l)
	if err != nil {
		return Persistents{}, err
	}

	w, err := webclients.Init(conf, l)
	if err != nil {
		return Persistents{}, err
	}

	return Persistents{
		Repositories: repo,
		Webclients:   w,
	}, nil
}
