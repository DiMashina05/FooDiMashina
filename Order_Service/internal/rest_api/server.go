package restapi

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
)

func NewHandler(logger *slog.Logger) *chi.Mux {
	handler := &Handler{logger: logger}

	r := chi.NewRouter()
	r.Use(loggingMiddleware(logger))

	r.Get("/order/{id}", handler.getOrders)
	r.Post("/order", handler.postOrders)
	r.Get("/health", handler.getHealth)
	return r
}
