package repository

import (
	"context"
	"errors"
	"id-topup-saas/backend/config"
	"id-topup-saas/backend/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type userPostgresRepository struct{}

func NewUserPostgresRepository() domain.UserRepository {
	return &userPostgresRepository{}
}

func (r *userPostgresRepository) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (id, tenant_id, name, email, password, role, balance) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := config.DB.Exec(ctx, query, user.ID, user.TenantID, user.Name, user.Email, user.Password, user.Role, user.Balance)
	return err
}

func (r *userPostgresRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, tenant_id, name, email, password, role, balance FROM users WHERE email = $1`
	var user domain.User
	err := config.DB.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.TenantID, &user.Name, &user.Email, &user.Password, &user.Role, &user.Balance,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userPostgresRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	query := `SELECT id, tenant_id, name, email, password, role, balance FROM users WHERE id = $1`
	var user domain.User
	err := config.DB.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.TenantID, &user.Name, &user.Email, &user.Password, &user.Role, &user.Balance,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userPostgresRepository) UpdateBalance(ctx context.Context, id uuid.UUID, delta float64) error {
	query := `UPDATE users SET balance = balance + $2, updated_at = NOW() WHERE id = $1`
	result, err := config.DB.Exec(ctx, query, id, delta)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("user not found")
	}
	return nil
}
