package repository

import (
	"context"
	"errors"
	"id-topup-saas/backend/internal/domain"
	"sync"

	"github.com/google/uuid"
)

type depositMemoryRepository struct {
	mu       sync.RWMutex
	deposits map[uuid.UUID]*domain.Deposit
}

func NewDepositMemoryRepository() domain.DepositRepository {
	return &depositMemoryRepository{
		deposits: make(map[uuid.UUID]*domain.Deposit),
	}
}

func (r *depositMemoryRepository) Create(ctx context.Context, deposit *domain.Deposit) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.deposits[deposit.ID] = deposit
	return nil
}

func (r *depositMemoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Deposit, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	deposit, ok := r.deposits[id]
	if !ok {
		return nil, errors.New("deposit not found")
	}
	return deposit, nil
}

func (r *depositMemoryRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	deposit, ok := r.deposits[id]
	if !ok {
		return errors.New("deposit not found")
	}

	deposit.Status = status
	return nil
}
