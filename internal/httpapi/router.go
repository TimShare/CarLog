package httpapi

import (
	"net/http"

	"carlog/internal/httpapi/cars"
	"carlog/internal/httpapi/system"
	"carlog/internal/service/car"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(pool *pgxpool.Pool, carService *car.Service) http.Handler {
	router := chi.NewRouter()
	router.Mount("/", system.NewHandler(pool).Routes())
	router.Mount("/cars", cars.NewHandler(carService).Routes())

	return router
}
