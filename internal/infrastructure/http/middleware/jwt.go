package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// errorResponse es una versión mínima local del formato estándar de respuesta.
type errorResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func newErrorResponse(message string) errorResponse {
	return errorResponse{
		Code:    "001",
		Message: message,
		Data:    nil,
	}
}

// JWTMiddleware valida el JWT en Authorization: Bearer <token>.
// Si JWT_SECRET no está definido, el middleware deja pasar todo (modo dev).
func JWTMiddleware() fiber.Handler {
	secret := os.Getenv("JWT_SECRET")

	return func(c *fiber.Ctx) error {
		// Si no hay secret configurado, NO exigimos JWT (útil para desarrollo local)
		if secret == "" {
			return c.Next()
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			res := newErrorResponse("missing Authorization header")
			return c.Status(http.StatusUnauthorized).JSON(res)
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			res := newErrorResponse("invalid Authorization header format")
			return c.Status(http.StatusUnauthorized).JSON(res)
		}

		tokenStr := parts[1]

		// Parsear y validar el token
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			// Validar que sea método HMAC (HS256, etc.)
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			res := newErrorResponse("invalid or expired token")
			return c.Status(http.StatusUnauthorized).JSON(res)
		}

		// Opcional: guardar el token/claims en el contexto
		c.Locals("jwt", token)

		return c.Next()
	}
}
