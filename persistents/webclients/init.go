package webclients

import (
	"github.com/naufalfmm/cryptocurrency-price-api/persistents/webclients/coincap"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/config"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

type Webclients struct {
	Coincap coincap.Coincap
}

func Init(conf *config.EnvConfig, log logger.Logger) (Webclients, error) {
	c, err := coincap.Init(conf.CoincapBasePath, log)
	if err != nil {
		return Webclients{}, err
	}

	return Webclients{
		Coincap: c,
	}, nil
}
