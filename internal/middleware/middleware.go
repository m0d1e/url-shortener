package middleware

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Common() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middleware.URLFormat,
	}
}
