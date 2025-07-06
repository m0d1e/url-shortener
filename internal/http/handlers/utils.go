package handlers

import (
	"encoding/json"
	"net/http"
	"url_shortener/internal/http/dto"
)

func sendResponse(w http.ResponseWriter, status int, resp dto.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}
