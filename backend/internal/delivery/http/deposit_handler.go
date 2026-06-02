

package http

import (
	"id-topup-saas/backend/internal/domain"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type DepositHandler struct {
	depositUsecase domain.DepositUsecase
}
func NewDepositHandler(router fiber.Router, usecase domain.DepositUsecase) {
	handler := &DepositHandler{
		depositUsecase: usecase,
	}

	router.Post("/deposit", JWTMiddleware(), handler.RequestDeposit)
}

func (h *DepositHandler) RequestDeposit(c fiber.Ctx) error {
	type DepositRequest struct {
		Amount float64 `json:"amount"`
	}

	var req DepositRequest
	if err := c.Bind().Body(&req); err != nil || req.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid amount"})
	}

	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	userID, _ := uuid.Parse(userIDStr)
	deposit, err := h.depositUsecase.RequestDeposit(c.Context(), userID, req.Amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "deposit request created",
		"data":    deposit,
	})
}
