package cars

import (
	"net/http"

	"carlog/internal/service/car"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *car.Service
}

func NewHandler(service *car.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes() http.Handler {
	router := chi.NewRouter()
	router.Post("/", h.create)

	return router
}
