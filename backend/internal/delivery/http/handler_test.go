package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"id-topup-saas/backend/internal/repository"
	"id-topup-saas/backend/internal/usecase"

	"github.com/gofiber/fiber/v3"
)

func TestRegisterEndpoint(t *testing.T) {
	app := fiber.New()
	api := app.Group("/api")

	userRepo := repository.NewUserMemoryRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	NewUserHandler(api, userUsecase)

	payload := []byte(`{"tenant_id":"00000000-0000-0000-0000-000000000001","name":"Tester","email":"tester@example.com","password":"secret123"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}
