package dialers

import (
	"github.com/naufalfmm/cryptocurrency-price-api/infrastructures/wslistener/listeners"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/config"
	"github.com/naufalfmm/cryptocurrency-price-api/usecases"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/listener/ws"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

type Dialers struct {
	Listeners listeners.Listeners
}

func Init(u usecases.Usecases, log logger.Logger) (Dialers, error) {
	l, err := listeners.Init(u, log)
	if err != nil {
		return Dialers{}, err
	}

	return Dialers{
		Listeners: l,
	}, nil
}

func (d Dialers) Register(we *ws.Engine, conf *config.EnvConfig) {
	if conf.CoincapPriceSyncMode {
		we.Dial("wss://ws.coincap.io/prices?assets=ALL", d.Listeners.Coins.SyncPrice)
	}
}
