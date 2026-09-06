package system

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type databasePinger interface {
	Ping(context.Context) error
}

type Handler struct {
	database databasePinger
}

func NewHandler(database databasePinger) *Handler {
	return &Handler{database: database}
}

func (h *Handler) Routes() http.Handler {
	router := chi.NewRouter()
	router.Get("/ping", h.ping)
	router.Get("/health", h.health)

	return router
}
