package usecase

import (
	"context"
	"errors"
	"id-topup-saas/backend/internal/domain"
	"time"

	"github.com/google/uuid"
)

type transactionUsecase struct {
	transactionRepo domain.TransactionRepository
	userRepo        domain.UserRepository
}

func NewTransactionUsecase(repo domain.TransactionRepository, userRepo domain.UserRepository) domain.TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: repo,
		userRepo:        userRepo,
	}
}

func (u *transactionUsecase) CreateTransaction(ctx context.Context, userID uuid.UUID, productID string, amount float64) (*domain.Transaction, error) {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user.Balance < amount {
		return nil, errors.New("insufficient balance")
	}

	if err := u.userRepo.UpdateBalance(ctx, userID, -amount); err != nil {
		return nil, err
	}

	newTransaction := &domain.Transaction{
		ID:        uuid.New(),
		UserID:    userID,
		ProductID: productID,
		Amount:    amount,
		Status:    "success",
		CreatedAt: time.Now(),
	}

	err = u.transactionRepo.Create(ctx, newTransaction)
	if err != nil {
		return nil, err
	}

	return newTransaction, nil
}
