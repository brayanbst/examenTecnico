package http

import (
	"github.com/gofiber/fiber/v2"

	mw "github.com/brayanbst/matrix-service-go/internal/infrastructure/http/middleware"
)

// RegisterRoutes registra las rutas HTTP de la API.
func RegisterRoutes(app *fiber.App, qrHandler *QRHandler, authHandler *AuthHandler) {
	// Rutas de autenticación (SIN JWT)
	auth := app.Group("/auth")
	auth.Post("/login", authHandler.PostLogin)

	// Grupo /api protegido por JWT
	api := app.Group("/api", mw.JWTMiddleware())

	// POST /api/qr
	api.Post("/qr", qrHandler.PostQR)

	// POST /api/qr-and-stats
	api.Post("/qr-and-stats", qrHandler.PostQRAndStats)
}
