package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	ProductID string    `json:"product_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type TransactionRepository interface {
	Create(ctx context.Context, transaction *Transaction) error
}

type TransactionUsecase interface {
	CreateTransaction(ctx context.Context, userID uuid.UUID, productID string, amount float64) (*Transaction, error)
}
