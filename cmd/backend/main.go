package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BarTar213/go-template/api"
	"github.com/BarTar213/go-template/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfgPath := flag.String("cfgPath", "config.yml", "path to config file")
	flag.Parse()

	conf := config.NewConfig(*cfgPath)
	log.Infof("%+v", conf)

	if conf.Api.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	a := api.NewApi(
		api.WithConfig(conf),
	)

	go a.Run()
	log.Info("started app")

	shutDownSignal := make(chan os.Signal)
	signal.Notify(shutDownSignal, syscall.SIGINT, syscall.SIGTERM)

	<-shutDownSignal
	a.Shutdown(5 * time.Second)
	log.Info("exited from app")
}
