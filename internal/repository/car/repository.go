package car

import (
	"context"
	_ "embed"
	"errors"
	"fmt"

	domain "carlog/internal/service/car"
	"carlog/internal/sqltemplate"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

//go:embed sql/create.sql.j2
var createQuerySource string

var createQuery = sqltemplate.MustParse("car/create", createQuerySource)

type queryer interface {
	QueryRow(context.Context, string, ...any) pgx.Row
}

type Repository struct {
	database queryer
}

func New(database queryer) *Repository {
	return &Repository{database: database}
}

func (r *Repository) Create(ctx context.Context, params domain.CreateCarParams) (created domain.Car, err error) {
	hasVIN := params.VIN != nil
	query, err := createQuery.Render(map[string]any{"has_vin": hasVIN})
	if err != nil {
		return domain.Car{}, err
	}

	arguments := []any{params.ModelID}
	if hasVIN {
		arguments = append(arguments, params.VIN)
	}
	arguments = append(arguments, params.Year, params.CurrentMileage)

	err = r.database.QueryRow(ctx, query, arguments...).Scan(
		&created.ID,
		&created.VIN,
		&created.Year,
		&created.CurrentMileage,
		&created.CreatedAt,
		&created.UpdatedAt,
	)
	if err != nil {
		var postgresError *pgconn.PgError
		if errors.As(err, &postgresError) && postgresError.ConstraintName == "cars_vin_unique_idx" {
			return domain.Car{}, domain.ErrVINExists
		}
		return domain.Car{}, fmt.Errorf("insert car: %w", err)
	}

	return created, nil
}
