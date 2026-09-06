package car

import (
	"context"
	"errors"
)

var ErrVINExists = errors.New("car with this VIN already exists")

type Repository interface {
	Create(context.Context, CreateCarParams) (Car, error)
}

type MakeRepository interface {
	Create(context.Context, string) (Make, error)
}

type ModelRepository interface {
	Create(context.Context, int64, string) (Model, error)
}

type Repositories struct {
	Makes  MakeRepository
	Models ModelRepository
	Cars   Repository
}

type TransactionManager interface {
	RunInTransaction(context.Context, func(Repositories) error) error
}
