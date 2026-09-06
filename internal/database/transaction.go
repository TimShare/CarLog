package database

import (
	"context"
	"fmt"

	carrepository "carlog/internal/repository/car"
	makerepository "carlog/internal/repository/make"
	modelrepository "carlog/internal/repository/model"
	"carlog/internal/service/car"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionManager struct {
	pool *pgxpool.Pool
}

func NewTransactionManager(pool *pgxpool.Pool) *TransactionManager {
	return &TransactionManager{pool: pool}
}

func (m *TransactionManager) RunInTransaction(ctx context.Context, fn func(car.Repositories) error) error {
	tx, err := m.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	repositories := car.Repositories{
		Makes:  makerepository.New(tx),
		Models: modelrepository.New(tx),
		Cars:   carrepository.New(tx),
	}
	if err := fn(repositories); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	return nil
}
