package make

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

func (r *Repository) Create(ctx context.Context, name string) (car.Make, error) {
	var make car.Make
	err := r.database.QueryRow(ctx, insertQuery, name).Scan(
		&make.ID,
		&make.Name,
		&make.CountryCode,
		&make.LogoURL,
	)
	if err != nil {
		return car.Make{}, fmt.Errorf("insert car make: %w", err)
	}

	return make, nil
}
