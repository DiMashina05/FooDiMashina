package restapi

import (
	"net/http"

	"github.com/DiMashina05/FooDiMashina/User_Service/internal/repository"
)

type Handler struct{
	repo repository.Repository
}

func (h *Handler) postUsers(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) postItems(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) deleteItems(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) getCart(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) getHealth(w http.ResponseWriter, r *http.Request) {}
