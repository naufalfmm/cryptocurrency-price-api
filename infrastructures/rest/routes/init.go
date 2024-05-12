package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/cryptocurrency-price-api/infrastructures/rest/controllers"
	"github.com/naufalfmm/cryptocurrency-price-api/middlewares"
	"github.com/naufalfmm/cryptocurrency-price-api/usecases"
)

type Routes struct {
	Controllers controllers.Controllers
	Middlewares middlewares.Middlewares
}

func Init(usec usecases.Usecases, middl middlewares.Middlewares) (Routes, error) {
	c, err := controllers.Init(usec)
	if err != nil {
		return Routes{}, err
	}

	return Routes{
		Controllers: c,
		Middlewares: middl,
	}, nil
}

func (r Routes) Register(ge *gin.Engine) {
	ge.RedirectFixedPath = true

	auth := ge.Group("/auth")
	auth.POST("/signin", r.Controllers.Auth.SignIn)
	auth.POST("/signup", r.Controllers.Auth.SignUp)

	userCoin := ge.Group("/user-coins")
	userCoin.POST("/track", r.Middlewares.VerifyToken(), r.Controllers.UserCoins.TrackCoin)
	userCoin.POST("/untrack", r.Middlewares.VerifyToken(), r.Controllers.UserCoins.UntrackCoin)
	userCoin.GET("", r.Middlewares.VerifyToken(), r.Controllers.UserCoins.GetAllTrack)
}
