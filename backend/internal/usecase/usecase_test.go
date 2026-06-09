package usecase

import (
	"context"
	"testing"

	"id-topup-saas/backend/internal/repository"

	"github.com/google/uuid"
)

func TestUserUsecase_RegisterLoginAndGetProfile(t *testing.T) {
	ctx := context.Background()
	userRepo := repository.NewUserMemoryRepository()
	uc := NewUserUsecase(userRepo)

	tenantID := uuid.New()
	user, err := uc.Register(ctx, tenantID, "Alice", "alice@example.com", "secret123")
	if err != nil {
		t.Fatalf("Register() error = %v", err)
	}
	if user.ID == uuid.Nil {
		t.Fatal("Register() returned empty user ID")
	}
	if user.TenantID != tenantID {
		t.Fatalf("Register() tenant ID = %v, want %v", user.TenantID, tenantID)
	}

	token, err := uc.Login(ctx, "alice@example.com", "secret123")
	if err != nil {
		t.Fatalf("Login() error = %v", err)
	}
	if token == "" {
		t.Fatal("Login() returned empty token")
	}

	profile, err := uc.GetProfile(ctx, user.ID)
	if err != nil {
		t.Fatalf("GetProfile() error = %v", err)
	}
	if profile.Email != "alice@example.com" {
		t.Fatalf("GetProfile() email = %q, want %q", profile.Email, "alice@example.com")
	}
}

func TestDepositAndTransactionUsecases_UpdateBalance(t *testing.T) {
	ctx := context.Background()
	userRepo := repository.NewUserMemoryRepository()
	depositRepo := repository.NewDepositMemoryRepository()
	transactionRepo := repository.NewTransactionMemoryRepository()

	userUC := NewUserUsecase(userRepo)
	user, err := userUC.Register(ctx, uuid.New(), "Bob", "bob@example.com", "secret123")
	if err != nil {
		t.Fatalf("Register() error = %v", err)
	}

	depositUC := NewDepositUsecase(depositRepo, userRepo)
	deposit, err := depositUC.RequestDeposit(ctx, user.ID, 50)
	if err != nil {
		t.Fatalf("RequestDeposit() error = %v", err)
	}
	if deposit.Amount != 50 {
		t.Fatalf("RequestDeposit() amount = %v, want 50", deposit.Amount)
	}

	profile, err := userRepo.GetByID(ctx, user.ID)
	if err != nil {
		t.Fatalf("GetByID() error = %v", err)
	}
	if profile.Balance != 50 {
		t.Fatalf("balance after deposit = %v, want 50", profile.Balance)
	}

	transactionUC := NewTransactionUsecase(transactionRepo, userRepo)
	transaction, err := transactionUC.CreateTransaction(ctx, user.ID, "product-1", 20)
	if err != nil {
		t.Fatalf("CreateTransaction() error = %v", err)
	}
	if transaction.ProductID != "product-1" {
		t.Fatalf("CreateTransaction() product ID = %q, want %q", transaction.ProductID, "product-1")
	}

	profile, err = userRepo.GetByID(ctx, user.ID)
	if err != nil {
		t.Fatalf("GetByID() after transaction error = %v", err)
	}
	if profile.Balance != 30 {
		t.Fatalf("balance after transaction = %v, want 30", profile.Balance)
	}
}
