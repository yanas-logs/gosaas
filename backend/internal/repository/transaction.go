package repository

import (
	"context"
	"id-topup-saas/backend/internal/domain"
	"sync"
)

type transactionMemoryRepository struct {
	mu           sync.RWMutex
	transactions []*domain.Transaction
}

func NewTransactionMemoryRepository() domain.TransactionRepository {
	return &transactionMemoryRepository{
		transactions: make([]*domain.Transaction, 0),
	}
}

func (r *transactionMemoryRepository) Create(ctx context.Context, transaction *domain.Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.transactions = append(r.transactions, transaction)
	return nil
}
