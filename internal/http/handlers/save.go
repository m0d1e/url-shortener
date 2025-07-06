package handlers

import (
	"encoding/json"
	"strings"
	"url_shortener/internal/slErr"
	
	//"log/slog"
	"net/http"
	"url_shortener/internal/http/dto"
)

func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
	var req dto.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.Error("failed to decode request body", slErr.Err(err))
		sendResponse(w, http.StatusBadRequest, dto.Response{
			Status: "fail",
			Error: "failed to decode request",
		})
		return
	}
	h.log.Info("request body decoded")
	
	alias, err := h.service.Save(req)
	if err != nil {
		if strings.Contains(err.Error(), "validation error") {
			h.log.Error("validation error", slErr.Err(err))
			sendResponse(w, http.StatusBadRequest, dto.Response{
				Status: "fail",
				Error: err.Error(),
			})
			return
		}
		
		h.log.Error("internal server error", slErr.Err(err))
		sendResponse(w, http.StatusInternalServerError, dto.Response{
			Status: "error",
			Error: "internal server error",
		})
		return
	}
	h.log.Info("url added")
	sendResponse(w, http.StatusCreated, dto.Response{
		Status: "ok",
		Alias: alias,
	})
}
