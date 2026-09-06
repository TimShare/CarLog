package car

import (
	"context"
	"fmt"
	"strings"
)

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

type Service struct {
	transactions TransactionManager
}

func NewService(transactions TransactionManager) *Service {
	return &Service{transactions: transactions}
}

func (s *Service) Create(ctx context.Context, params CreateParams) (Car, error) {
	params.Make = strings.TrimSpace(params.Make)
	params.Model = strings.TrimSpace(params.Model)

	if params.Make == "" {
		return Car{}, ValidationError{Message: "make is required"}
	}
	if params.Model == "" {
		return Car{}, ValidationError{Message: "model is required"}
	}
	if params.VIN != nil {
		normalizedVIN := strings.ToUpper(strings.TrimSpace(*params.VIN))
		if normalizedVIN == "" {
			return Car{}, ValidationError{Message: "vin must not be empty"}
		}
		if len(normalizedVIN) > 17 {
			return Car{}, ValidationError{Message: "vin must not be longer than 17 characters"}
		}
		params.VIN = &normalizedVIN
	}
	if params.Year < 1886 || params.Year > 2100 {
		return Car{}, ValidationError{Message: "year must be between 1886 and 2100"}
	}
	if params.CurrentMileage < 0 {
		return Car{}, ValidationError{Message: "current_mileage must not be negative"}
	}

	var created Car
	err := s.transactions.RunInTransaction(ctx, func(repositories Repositories) error {
		make, err := repositories.Makes.Create(ctx, params.Make)
		if err != nil {
			return fmt.Errorf("create make: %w", err)
		}

		model, err := repositories.Models.Create(ctx, make.ID, params.Model)
		if err != nil {
			return fmt.Errorf("create model: %w", err)
		}

		created, err = repositories.Cars.Create(ctx, CreateCarParams{
			ModelID:        model.ID,
			VIN:            params.VIN,
			Year:           params.Year,
			CurrentMileage: params.CurrentMileage,
		})
		if err != nil {
			return err
		}

		created.Make = make.Name
		created.Model = model.Name
		return nil
	})
	if err != nil {
		return Car{}, fmt.Errorf("create car: %w", err)
	}

	return created, nil
}
