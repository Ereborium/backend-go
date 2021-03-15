package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BarTar213/go-template/api"
	"github.com/BarTar213/go-template/config"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.NewConfig("app_name.yml")
	logger := log.New(os.Stdout, "", log.LstdFlags)

	logger.Printf("%+v\n", conf)

	if conf.Api.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	a := api.NewApi(
		api.WithConfig(conf),
		api.WithLogger(logger),
	)

	go a.Run()
	logger.Print("started app")

	shutDownSignal := make(chan os.Signal)
	signal.Notify(shutDownSignal, syscall.SIGINT, syscall.SIGTERM)

	<-shutDownSignal
	logger.Print("exited from app")
}
