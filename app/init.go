package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/naufalfmm/cryptocurrency-price-api/infrastructures"
	"github.com/naufalfmm/cryptocurrency-price-api/middlewares"
	"github.com/naufalfmm/cryptocurrency-price-api/persistents"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/config"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/db"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/jwtoken"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/log"
	"github.com/naufalfmm/cryptocurrency-price-api/resources/validator"
	"github.com/naufalfmm/cryptocurrency-price-api/usecases"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/listener/ws"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/password/bcrypt"

	validatorUtils "github.com/naufalfmm/cryptocurrency-price-api/utils/validator"
)

type App struct {
	ge *gin.Engine
	we *ws.Engine

	middlewares middlewares.Middlewares
	validator   validatorUtils.Validator

	conf   *config.EnvConfig
	logger logger.Logger
}

func Init() App {
	ge := gin.New()
	we := ws.New()

	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	log, err := log.NewLogger(conf)
	if err != nil {
		panic(err)
	}

	pass, err := bcrypt.NewBcrypt(bcrypt.WithCost(conf.BcryptCost))
	if err != nil {
		panic(err)
	}

	jwtok, err := jwtoken.NewJWT(conf)
	if err != nil {
		panic(err)
	}

	orm, err := db.NewSqlite(conf, log)
	if err != nil {
		panic(err)
	}

	persist, err := persistents.Init(orm, log, conf)
	if err != nil {
		panic(err)
	}

	usec, err := usecases.Init(persist, pass, jwtok, log, orm, conf)
	if err != nil {
		panic(err)
	}

	validator, err := validator.NewValidator()
	if err != nil {
		panic(err)
	}

	middl, err := middlewares.Init(usec, jwtok)
	if err != nil {
		panic(err)
	}

	infr, err := infrastructures.Init(usec, middl, log)
	if err != nil {
		panic(err)
	}

	infr.Register(ge, we, conf)

	return App{
		ge: ge,
		we: we,

		middlewares: middl,
		validator:   validator,

		conf:   conf,
		logger: log,
	}
}

func (app App) Run() {
	binding.Validator = app.validator

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", app.conf.Port),
		Handler: app.ge,
	}

	go httpServer.ListenAndServe()
	go app.we.Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(
		sc,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	<-sc

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctxShutDown); err != nil {
		panic(err)
	}
}
