package main

import (
	"log"
	"os"

	"id-topup-saas/backend/internal/delivery/http"
	"id-topup-saas/backend/internal/repository"
	"id-topup-saas/backend/internal/usecase"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	apiGroup := app.Group("/api")

	userRepo := repository.NewUserMemoryRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	http.NewUserHandler(apiGroup, userUsecase)

	depositRepo := repository.NewDepositMemoryRepository()
	depositUsecase := usecase.NewDepositUsecase(depositRepo, userRepo)
	http.NewDepositHandler(apiGroup, depositUsecase)

	txRepo := repository.NewTransactionMemoryRepository()
	txUsecase := usecase.NewTransactionUsecase(txRepo, userRepo)
	http.NewTransactionHandler(apiGroup, txUsecase)

	app.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("Pong! Backend ID Topup SaaS Ready")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
