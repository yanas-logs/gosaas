package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Deposit struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type DepositRepository interface {
	Create(ctx context.Context, deposit *Deposit) error
	GetByID(ctx context.Context, id uuid.UUID) (*Deposit, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
}

type DepositUsecase interface {
	RequestDeposit(ctx context.Context, userID uuid.UUID, amount float64) (*Deposit, error)
	GetDepositByID(ctx context.Context, id uuid.UUID) (*Deposit, error)
}
