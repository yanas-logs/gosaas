package repository

import (
	"context"
	"errors"
	"sync"
	"time"

	"id-topup-saas/backend/internal/domain"

	"github.com/google/uuid"
)

type userMemoryRepository struct {
	mu    sync.RWMutex
	users map[uuid.UUID]*domain.User
}

func NewUserMemoryRepository() domain.UserRepository {
	return &userMemoryRepository{
		users: make(map[uuid.UUID]*domain.User),
	}
}

func (r *userMemoryRepository) Create(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return nil
}

func (r *userMemoryRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, user := range r.users {
		if user.Email == email {
			return cloneUser(user), nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *userMemoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return cloneUser(user), nil
}

func cloneUser(user *domain.User) *domain.User {
	if user == nil {
		return nil
	}

	clone := *user
	return &clone
}

func (r *userMemoryRepository) UpdateBalance(ctx context.Context, id uuid.UUID, delta float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return errors.New("user not found")
	}

	user.Balance += delta
	user.UpdatedAt = time.Now()
	return nil
}
