package restapi

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
)

func NewHandler(logger *slog.Logger) *chi.Mux {
	handler := &Handler{}

	r := chi.NewRouter()
	r.Use(loggingMiddleware(logger))

	r.Get("/health", handler.getHealth)
	r.Get("/cart", handler.getCart)
	r.Get("/user", handler.getUser)

	r.Post("/cart", handler.postCart)
	r.Post("/goods", handler.postGoods)
	r.Post("/user", handler.postUser)

	return r
}
