package restapi

import (
	"net/http"
)

type Handler struct{}

func (h *Handler) postUser(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) postCart(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) postGoods(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) getCart(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) getHealth(w http.ResponseWriter, r *http.Request) {}
