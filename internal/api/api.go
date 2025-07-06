package api

import (
	"github.com/go-chi/chi/v5"
	middlewareog "github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"url_shortener/internal/config"
	"url_shortener/internal/http/handlers"
	"url_shortener/internal/middleware"
	"url_shortener/service"
)

type API struct {
	router  *chi.Mux
	handler *handlers.Handler
	cfg     *config.Config
}

func NewAPI(r *chi.Mux, repo service.Repository, logger *slog.Logger, config *config.Config) *API {
	for _, m := range middleware.Common() {
		r.Use(m)
	}
	
	svc := service.NewURLService(repo, logger)
	handler := handlers.NewHandler(svc, logger)
	
	return &API{
		router:  r,
		handler: handler,
		cfg:     config,
	}
}

func (api *API) Handle() {
	api.router.Get("/{alias}", api.handler.Redirect)
	api.router.Route("/url", func(r chi.Router) {
		r.Use(middlewareog.BasicAuth("url-shortener", map[string]string{
			api.cfg.User: api.cfg.Password,
		}))
		
		r.Post("/", api.handler.Save)
		r.Delete("/{alias}", api.handler.Delete)
	})
}