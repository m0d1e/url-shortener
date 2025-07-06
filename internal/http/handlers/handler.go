package handlers

import (
	"log/slog"
	"url_shortener/service"
)

type Handler struct {
	service *service.URLService
	log     *slog.Logger
}

func NewHandler(svc *service.URLService, logger *slog.Logger) *Handler {
	return &Handler{
		service: svc,
		log:     logger,
	}
}
