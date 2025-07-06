package handlers

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"strings"
	"url_shortener/internal/http/dto"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "alias")
	if alias == "" {
		h.log.Info("alias empty")
		sendResponse(w, http.StatusBadRequest, dto.Response{
			Status: "fail",
			Error: "not found alias",
		})
		
		return
	}
	
	err := h.service.Delete(alias)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.log.Info("alias not found")
			sendResponse(w, http.StatusNotFound, dto.Response{
				Status: "fail",
				Error: err.Error(),
			})
			return
		}
		
		h.log.Error("failed to delete url", slog.String("alias", alias))
		sendResponse(w, http.StatusInternalServerError, dto.Response{
			Status: "fail",
			Error: "internal error",
		})
		return
	}
	
	h.log.Info("deleted url", slog.String("alias", alias))
	sendResponse(w, http.StatusOK, dto.Response{
		Status: "ok",
	})
}
