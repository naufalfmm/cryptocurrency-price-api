package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/cryptocurrency-price-api/usecases"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/token/jwt"
)

type (
	Middlewares interface {
		ImplementCors() gin.HandlerFunc
		VerifyToken() gin.HandlerFunc
		PanicRecover() gin.HandlerFunc
	}

	middlewares struct {
		usecases usecases.Usecases
		jwt      jwt.JWT
	}
)

func Init(usec usecases.Usecases, jwt jwt.JWT) (Middlewares, error) {
	return &middlewares{
		usecases: usec,
		jwt:      jwt,
	}, nil
}
