package http

import (
	"id-topup-saas/backend/internal/domain"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type TransactionHandler struct {
	transactionUsecase domain.TransactionUsecase
}

func NewTransactionHandler(router fiber.Router, usecase domain.TransactionUsecase) {
	handler := &TransactionHandler{
		transactionUsecase: usecase,
	}

	router.Post("/transactions", JWTMiddleware(), handler.CreateTransaction)
}

func (h *TransactionHandler) CreateTransaction(c fiber.Ctx) error {
	type TransactionRequest struct {
		ProductID string  `json:"product_id"`
		Amount    float64 `json:"amount"`
	}

	var req TransactionRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	userIDStr := c.Locals("user_id").(string)
	userID, _ := uuid.Parse(userIDStr)

	tx, err := h.transactionUsecase.CreateTransaction(c.Context(), userID, req.ProductID, req.Amount)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "transaction successful",
		"data":    tx,
	})
}
