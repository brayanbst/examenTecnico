package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	appqr "github.com/brayanbst/matrix-service-go/internal/application/qr"
	httpinfra "github.com/brayanbst/matrix-service-go/internal/infrastructure/http"
	infraqr "github.com/brayanbst/matrix-service-go/internal/infrastructure/qr"
)

func main() {
	// Cargar .env (solo útil en local / desarrollo)
	_ = godotenv.Load()

	app := fiber.New()

	qrDecomposer := infraqr.NewGonumQRDecomposer()
	service := appqr.NewService(qrDecomposer)
	handler := httpinfra.NewQRHandler(service)

	httpinfra.RegisterRoutes(app, handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Go API listening on port %s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
