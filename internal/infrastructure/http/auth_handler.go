package http

import (
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AuthHandler maneja endpoints de autenticación (login).
type AuthHandler struct{}

// NewAuthHandler crea un nuevo AuthHandler.
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// DTO de login.
type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// DTO para el token (irá dentro de "data").
type loginData struct {
	Token string `json:"token"`
}

// PostLogin maneja POST /auth/login
func (h *AuthHandler) PostLogin(c *fiber.Ctx) error {
	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			NewErrorResponse("invalid JSON"),
		)
	}

	// ⚠️ Ejemplo simple: credenciales hardcodeadas solo para el challenge
	if req.Username != "admin" || req.Password != "secret" {
		return c.Status(http.StatusUnauthorized).JSON(
			NewErrorResponse("invalid credentials"),
		)
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return c.Status(http.StatusInternalServerError).JSON(
			NewErrorResponse("JWT_SECRET not configured"),
		)
	}

	// Claims del token
	claims := jwt.MapClaims{
		"sub":  req.Username,                               // subject (usuario)
		"role": "admin",                                    // ejemplo de rol
		"exp":  time.Now().Add(1 * time.Hour).Unix(),       // expira en 1 hora
		"iat":  time.Now().Unix(),                          // emitido en
	}

	// Crear token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar token
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			NewErrorResponse("could not sign token"),
		)
	}

	data := loginData{Token: signed}

	return c.Status(http.StatusOK).JSON(
		NewSuccessResponse("login successful", data),
	)
}
