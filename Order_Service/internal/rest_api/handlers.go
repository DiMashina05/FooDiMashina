package restapi

import (
	"net/http"
)

type Handler struct{}

func (h *Handler) getOrders(w http.ResponseWriter, r *http.Request) {
	//тут скоро что-
}

func (h *Handler) getHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
