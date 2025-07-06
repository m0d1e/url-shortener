package handlers

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"strings"
	"url_shortener/internal/http/dto"
)

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "alias")
	if alias == "" {
		h.log.Info("alias empty")
		sendResponse(w, http.StatusBadRequest, dto.Response{
			Status: "fail",
			Error: "not found alias",
		})
		
		return
	}
	
	url, err := h.service.Get(alias)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.log.Info("url not found")
			sendResponse(w, http.StatusNotFound, dto.Response{
				Status: "fail",
				Error: "url not found",
			})
			return
		}
		
		h.log.Error("failed to get url")
		sendResponse(w, http.StatusInternalServerError, dto.Response{
			Status: "fail",
			Error: "internal error",
		})
		return
	}
	
	h.log.Info("redirecting", slog.String("url", url))
	//sendResponse(w, http.StatusFound, dto.Response{
	//	Status: "ok",
	//	Url:   url,
	//})
	http.Redirect(w, r, url, http.StatusFound)
}
