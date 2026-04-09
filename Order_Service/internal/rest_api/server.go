package restapi

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
)

func NewHandler(logger *slog.Logger) *chi.Mux {
	handler := &Handler{}

	r := chi.NewRouter()
	r.Use(loggingMiddleware(logger))

	r.Get("/order/{id}", handler.getOrders)
	r.Get("/health", handler.getHealth)
	return r
}
