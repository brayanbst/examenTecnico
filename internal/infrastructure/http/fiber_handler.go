package http

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"

	appqr "github.com/brayanbst/matrix-service-go/internal/application/qr"
	"github.com/brayanbst/matrix-service-go/internal/domain/matrix"
)

// QRHandler maneja las peticiones HTTP relacionadas con QR.
type QRHandler struct {
	service *appqr.Service
}

func NewQRHandler(service *appqr.Service) *QRHandler {
	return &QRHandler{service: service}
}

// DTO de entrada HTTP.
type matrixRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

// qrData es el payload específico que va dentro de "data" para /qr.
type qrData struct {
	Q [][]float64 `json:"Q"`
	R [][]float64 `json:"R"`
}

// qrAndStatsData es el payload para /qr-and-stats.
type qrAndStatsData struct {
	Q     [][]float64  `json:"Q"`
	R     [][]float64  `json:"R"`
	Stats *appqr.Stats `json:"stats"`
}

// PostQR maneja POST /api/qr
func (h *QRHandler) PostQR(c *fiber.Ctx) error {
	var req matrixRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			NewErrorResponse("invalid JSON"),
		)
	}

	m, err := matrix.NewMatrix(req.Matrix)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			NewErrorResponse(err.Error()),
		)
	}

	qrResult, err := h.service.ComputeQR(context.Background(), m)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			NewErrorResponse("internal error computing QR"),
		)
	}

	data := qrData{
		Q: qrResult.Q,
		R: qrResult.R,
	}

	return c.Status(http.StatusOK).JSON(
		NewSuccessResponse("QR decomposition computed successfully", data),
	)
}

// PostQRAndStats maneja POST /api/qr-and-stats
func (h *QRHandler) PostQRAndStats(c *fiber.Ctx) error {
	var req matrixRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			NewErrorResponse("invalid JSON"),
		)
	}

	m, err := matrix.NewMatrix(req.Matrix)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			NewErrorResponse(err.Error()),
		)
	}

	// Obtenemos el header Authorization para reenviarlo a Node.
	authHeader := c.Get("Authorization")

	qrResult, stats, err := h.service.ComputeQRAndStats(context.Background(), m, authHeader)
	if err != nil {
		// Si el statsPort no está configurado
		if err == appqr.ErrStatsPortNotConfigured {
			return c.Status(http.StatusInternalServerError).JSON(
				NewErrorResponse("stats service not configured"),
			)
		}
		// Otros errores (fallo al llamar a Node, etc.)
		return c.Status(http.StatusInternalServerError).JSON(
			NewErrorResponse(err.Error()),
		)
	}

	data := qrAndStatsData{
		Q:     qrResult.Q,
		R:     qrResult.R,
		Stats: stats,
	}

	return c.Status(http.StatusOK).JSON(
		NewSuccessResponse("QR and statistics computed successfully", data),
	)
}
