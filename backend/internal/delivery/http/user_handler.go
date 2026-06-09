package http

import (
	"id-topup-saas/backend/internal/domain"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(router fiber.Router, usecase domain.UserUsecase) {
	handler := &UserHandler{
		userUsecase: usecase,
	}

	router.Post("/register", handler.Register)
	router.Post("/login", handler.Login)
	router.Post("/logout", handler.Logout)

	protected := router.Group("", JWTMiddleware())
	protected.Get("/profile", handler.GetProfile)
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	type RegisterRequest struct {
		TenantID string `json:"tenant_id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	tenantID, err := uuid.Parse(req.TenantID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid tenant_id"})
	}

	user, err := h.userUsecase.Register(c.Context(), tenantID, req.Name, req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "registration successful",
		"data":    user,
	})
}

func (h *UserHandler) Login(c fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	token, err := h.userUsecase.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login successful",
		"token":   token,
	})
}

func (h *UserHandler) Logout(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logout successful, please clear your token on client side",
	})
}

func (h *UserHandler) GetProfile(c fiber.Ctx) error {
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "identity not found"})
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id format"})
	}

	user, err := h.userUsecase.GetProfile(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "profile not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "profile retrieval successful",
		"data":    user,
	})
}
