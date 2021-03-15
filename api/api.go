package api

import (
	"log"
	"net/http"

	"github.com/BarTar213/go-template/config"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Port   string
	Router *gin.Engine
	Config *config.Config
	Logger *log.Logger
}

func WithConfig(conf *config.Config) func(a *Api) {
	return func(a *Api) {
		a.Config = conf
	}
}

func WithLogger(logger *log.Logger) func(a *Api) {
	return func(a *Api) {
		a.Logger = logger
	}
}

func NewApi(options ...func(api *Api)) *Api {
	a := &Api{
		Router: gin.Default(),
	}
	a.Router.Use(gin.Recovery())

	for _, option := range options {
		option(a)
	}

	a.Router.GET("/", a.health)

	return a
}

func (a *Api) Run() error {
	return a.Router.Run(a.Config.Api.Port)
}

func (a *Api) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "healthy")
}
