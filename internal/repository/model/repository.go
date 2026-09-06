package model

import (
	"context"
	_ "embed"
	"fmt"

	"carlog/internal/service/car"

	"github.com/jackc/pgx/v5"
)

//go:embed sql/insert.sql.j2
var insertQuery string

type queryer interface {
	QueryRow(context.Context, string, ...any) pgx.Row
}

type Repository struct {
	database queryer
}

func New(database queryer) *Repository {
	return &Repository{database: database}
}

func (r *Repository) Create(ctx context.Context, makeID int64, name string) (car.Model, error) {
	var model car.Model
	err := r.database.QueryRow(ctx, insertQuery, makeID, name).Scan(
		&model.ID,
		&model.MakeID,
		&model.Name,
		&model.ProductionStartYear,
		&model.ProductionEndYear,
	)
	if err != nil {
		return car.Model{}, fmt.Errorf("insert car model: %w", err)
	}

	return model, nil
}
