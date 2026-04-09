package restapi

import (
	"log/slog"

	"github.com/DiMashina05/FooDiMashina/User_Service/internal/repository"
	"github.com/go-chi/chi/v5"
)

func NewHandler(repo repository.Repository,logger *slog.Logger) *chi.Mux {
	handler := &Handler{repo: repo}

	r := chi.NewRouter()
	r.Use(loggingMiddleware(logger))

	r.Get("/health", handler.getHealth)
	r.Get("/users/{id}/cart", handler.getCart)
	r.Get("/users/{id}", handler.getUsers)

	r.Post("/users/{id}/cart/items", handler.postItems)
	r.Post("/users", handler.postUsers)

	r.Delete("/users/{id}/cart/items/{item_id}", handler.deleteItems)

	return r
}
