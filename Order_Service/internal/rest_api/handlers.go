package restapi

import (
	"net/http"
)

type Handler struct{}

func (h *Handler) postOrders(w http.ResponseWriter, r *http.Request) {

	//тут скоро что-то будет
}

func (h *Handler) getOrders(w http.ResponseWriter, r *http.Request) {
	//тут скоро что-то будет
}

func (h *Handler) getHealth(w http.ResponseWriter, r *http.Request) {
	//тут скоро что-то будет
}
