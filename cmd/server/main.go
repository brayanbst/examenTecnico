package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	appqr "github.com/brayanbst/matrix-service-go/internal/application/qr"
	httpinfra "github.com/brayanbst/matrix-service-go/internal/infrastructure/http"
	nodeclient "github.com/brayanbst/matrix-service-go/internal/infrastructure/nodeclient"
	infraqr "github.com/brayanbst/matrix-service-go/internal/infrastructure/qr"
)

func main() {
	app := fiber.New()

	// --- Infraestructura QR (Gonum) ---
	qrDecomposer := infraqr.NewGonumQRDecomposer()

	// --- Cliente HTTP hacia Node (stats) ---
	nodeBaseURL := os.Getenv("NODE_API_URL")
	if nodeBaseURL == "" {
		// Útil en local si Node corre en http://localhost:3000
		nodeBaseURL = "http://localhost:3000"
	}
	statsClient := nodeclient.NewHTTPStatsClient(nodeBaseURL)

	// --- Caso de uso (QR + Stats) ---
	service := appqr.NewService(qrDecomposer, statsClient)

	// --- Handlers HTTP ---
	qrHandler := httpinfra.NewQRHandler(service)
	authHandler := httpinfra.NewAuthHandler()

	// --- Rutas ---
	httpinfra.RegisterRoutes(app, qrHandler, authHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Go API listening on port %s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
