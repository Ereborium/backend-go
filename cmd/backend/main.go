package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/BarTar213/go-template/api"
	"github.com/BarTar213/go-template/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfgPath := flag.String("cfg", "config.yml", "path to config file")
	flag.Parse()

	conf, err := config.New(*cfgPath)
	if err != nil {
		log.Error(err)
		return
	}

	log.WithField("app", "backend-go")
	log.Infof("%+v", conf)

	a := api.NewApi(
		api.WithConfig(conf),
	)

	go a.Run()
	log.Info("started app")

	shutDownSignal := make(chan os.Signal)
	signal.Notify(shutDownSignal, syscall.SIGINT, syscall.SIGTERM)

	<-shutDownSignal
	a.Shutdown(conf.Api.Timeout)
	log.Info("exited from app")
}
