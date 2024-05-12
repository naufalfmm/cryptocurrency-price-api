package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/cryptocurrency-price-api/infrastructures/rest/routes"
	"github.com/naufalfmm/cryptocurrency-price-api/middlewares"
	"github.com/naufalfmm/cryptocurrency-price-api/usecases"
)

type Rest struct {
	Routes routes.Routes
}

func Init(usec usecases.Usecases, middl middlewares.Middlewares) (Rest, error) {
	rout, err := routes.Init(usec, middl)
	if err != nil {
		return Rest{}, err
	}

	return Rest{
		Routes: rout,
	}, nil
}

func (r *Rest) Register(ge *gin.Engine) {
	r.Routes.Register(ge)
}
