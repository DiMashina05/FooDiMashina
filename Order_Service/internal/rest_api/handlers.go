package restapi

import (
	"log/slog"
	"net/http"
)

type Handler struct {
	logger *slog.Logger
}

func (h *Handler) postOrders(w http.ResponseWriter, r *http.Request) {
	
	//
}

func (h *Handler) getOrders(w http.ResponseWriter, r *http.Request) {
	//
}

func (h *Handler) getHealth(w http.ResponseWriter, r *http.Request) {
	//
}
