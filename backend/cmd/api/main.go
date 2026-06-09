package main

import (
	"log"
	"os"

	"id-topup-saas/backend/config" // Pastikan import sesuai dengan nama module di go.mod Anda
	"id-topup-saas/backend/internal/delivery/http"
	"id-topup-saas/backend/internal/repository"
	"id-topup-saas/backend/internal/usecase"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	apiGroup := app.Group("/api")

	userRepo := repository.NewUserMemoryRepository()
	depositRepo := repository.NewDepositMemoryRepository()
	txRepo := repository.NewTransactionMemoryRepository()

	if err := config.ConnectDB(); err == nil {
		defer config.DB.Close()
		userRepo = repository.NewUserPostgresRepository()
	} else {
		log.Printf("PostgreSQL unavailable, using in-memory repositories: %v", err)
	}

	userUsecase := usecase.NewUserUsecase(userRepo)
	http.NewUserHandler(apiGroup, userUsecase)

	depositUsecase := usecase.NewDepositUsecase(depositRepo, userRepo)
	http.NewDepositHandler(apiGroup, depositUsecase)

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
