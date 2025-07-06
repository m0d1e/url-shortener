package main

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"os"
	"url_shortener/internal/api"
	"url_shortener/internal/config"
	"url_shortener/internal/slErr"
	"url_shortener/internal/storage"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func main() {
	// Init config
	cfg := config.LoadConfig()
	
	// Init logger
	log := setupLogger(cfg.Env)
	log.Info("start url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug on")
	
	// Init database
	db, err := storage.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Error("cannot innit db", slErr.Err(err))
		os.Exit(1)
	}
	
	// Init router
	router := chi.NewRouter()
	apiServer := api.NewAPI(router, db, log, cfg)
	apiServer.Handle()
	
	// Run server
	log.Info("starting HTTP server", slog.String("addr", cfg.HTTPServer.Address))
	if err := http.ListenAndServe(cfg.HTTPServer.Address, router); err != nil {
		log.Error("server exited with error", slErr.Err(err))
		os.Exit(1)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	
	return log
}
