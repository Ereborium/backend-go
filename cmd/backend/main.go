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

	cfg, err := config.New(*cfgPath)
	if err != nil {
		log.Error(err)
		return
	}

	log.WithField("app", "backend-go")
	log.Infof("%+v", cfg)

	a := api.NewApi(
		api.WithConfig(cfg),
	)

	go func() {
		err = a.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Info("started app")

	shutDownSignal := make(chan os.Signal, 1)
	signal.Notify(shutDownSignal, syscall.SIGINT, syscall.SIGTERM)

	<-shutDownSignal
	a.Shutdown(cfg.Api.Timeout)
	log.Info("exited from app")
}
