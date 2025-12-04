package http

import (
	"github.com/gofiber/fiber/v2"

	mw "github.com/brayanbst/matrix-service-go/internal/infrastructure/http/middleware"
)

func RegisterRoutes(app *fiber.App, qrHandler *QRHandler) {
	// --- Rutas públicas ---
	authHandler := NewAuthHandler()
	app.Post("/auth/login", authHandler.PostLogin)

	// --- Rutas protegidas con JWT ---
	api := app.Group("/api", mw.JWTMiddleware())
	api.Post("/qr", qrHandler.PostQR)
}
