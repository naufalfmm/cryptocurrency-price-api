package infrastructures

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/cryptocurrency-price-api/infrastructures/rest"
	"github.com/naufalfmm/cryptocurrency-price-api/infrastructures/wslistener"
	"github.com/naufalfmm/cryptocurrency-price-api/middlewares"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/config"
	"github.com/naufalfmm/cryptocurrency-price-api/usecases"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/listener/ws"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

type Infrastructures struct {
	Rest       rest.Rest
	WsListener wslistener.WsListener
}

func Init(usec usecases.Usecases, middl middlewares.Middlewares, log logger.Logger) (Infrastructures, error) {
	res, err := rest.Init(usec, middl)
	if err != nil {
		return Infrastructures{}, err
	}

	wsl, err := wslistener.Init(usec, log)
	if err != nil {
		return Infrastructures{}, err
	}

	return Infrastructures{
		Rest:       res,
		WsListener: wsl,
	}, nil
}

func (i *Infrastructures) Register(ge *gin.Engine, we *ws.Engine, conf *config.EnvConfig) {
	i.Rest.Register(ge)
	i.WsListener.Register(we, conf)
}
