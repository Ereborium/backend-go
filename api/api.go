package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/BarTar213/go-template/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
)

type Api struct {
	Port   string
	Router chi.Router
	server *http.Server
	Config *config.Config
}

func WithConfig(conf *config.Config) func(a *Api) {
	return func(a *Api) {
		a.Config = conf
	}
}

func NewApi(options ...func(api *Api)) *Api {
	a := &Api{
		Router: chi.NewRouter(),
	}
	a.Router.Use(middleware.Recoverer, middleware.Logger)

	for _, option := range options {
		option(a)
	}

	a.server = &http.Server{
		Addr:    a.Config.Api.Port,
		Handler: a.Router,
	}

	a.Router.Get("/", a.health)

	return a
}

func (a *Api) Run() error {
	return a.server.ListenAndServe()
}

func (a *Api) Shutdown(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		log.Errorf("server forced to shutdown: %s", err)
	}
}

func (a *Api) health(w http.ResponseWriter, r *http.Request) {
	respond(w, http.StatusOK, "healthy")
}

func respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")

	if err := enc.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
}
