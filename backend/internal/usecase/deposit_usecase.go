package usecase

import (
	"context"
	"id-topup-saas/backend/internal/domain"
	"time"

	"github.com/google/uuid"
)

type depositUsecase struct {
	depositRepo domain.DepositRepository
	userRepo    domain.UserRepository
}

func NewDepositUsecase(repo domain.DepositRepository, userRepo domain.UserRepository) domain.DepositUsecase {
	return &depositUsecase{
		depositRepo: repo,
		userRepo:    userRepo,
	}
}

func (u *depositUsecase) RequestDeposit(ctx context.Context, userID uuid.UUID, amount float64) (*domain.Deposit, error) {
	if _, err := u.userRepo.GetByID(ctx, userID); err != nil {
		return nil, err
	}

	err := u.userRepo.UpdateBalance(ctx, userID, amount)
	if err != nil {
		return nil, err
	}

	newDeposit := &domain.Deposit{
		ID:        uuid.New(),
		UserID:    userID,
		Amount:    amount,
		Status:    "success",
		CreatedAt: time.Now(),
	}

	err = u.depositRepo.Create(ctx, newDeposit)
	if err != nil {
		return nil, err
	}

	return newDeposit, nil
}

func (u *depositUsecase) GetDepositByID(ctx context.Context, id uuid.UUID) (*domain.Deposit, error) {
	return u.depositRepo.GetByID(ctx, id)
}
