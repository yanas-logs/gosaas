package repository

import (
	"context"
	"testing"

	"id-topup-saas/backend/internal/domain"

	"github.com/google/uuid"
)

func TestUserMemoryRepository_GetByID_ReturnsSafeCopy(t *testing.T) {
	ctx := context.Background()
	repo := NewUserMemoryRepository().(*userMemoryRepository)

	user := &domain.User{
		ID:       uuid.New(),
		TenantID: uuid.New(),
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "hashed",
		Balance:  100,
		Role:     "reseller",
	}

	if err := repo.Create(ctx, user); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	fetched, err := repo.GetByID(ctx, user.ID)
	if err != nil {
		t.Fatalf("GetByID() error = %v", err)
	}

	fetched.Balance = 999

	stored, err := repo.GetByID(ctx, user.ID)
	if err != nil {
		t.Fatalf("GetByID() second call error = %v", err)
	}

	if stored.Balance != 100 {
		t.Fatalf("stored balance = %v, want 100", stored.Balance)
	}
}
