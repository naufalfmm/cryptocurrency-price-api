package wslistener

import (
	"github.com/naufalfmm/cryptocurrency-price-api/infrastructures/wslistener/dialers"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/config"
	"github.com/naufalfmm/cryptocurrency-price-api/usecases"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/listener/ws"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

type WsListener struct {
	Dialers dialers.Dialers
}

func Init(usec usecases.Usecases, log logger.Logger) (WsListener, error) {
	dial, err := dialers.Init(usec, log)
	if err != nil {
		return WsListener{}, err
	}

	return WsListener{
		Dialers: dial,
	}, nil
}

func (r *WsListener) Register(we *ws.Engine, conf *config.EnvConfig) {
	r.Dialers.Register(we, conf)
}
