package server

import (
	"crypto/subtle"
	"strings"

	"github.com/gofiber/fiber/v3"
)

// APIKeyAuth protects all application routes except operational endpoints.
// Comparisons use constant time and the configured secret is never logged.
func APIKeyAuth(expected string) fiber.Handler {
	return func(c fiber.Ctx) error {
		path := c.Path()
		if c.Method() == fiber.MethodOptions || path == "/health" || path == "/docs" || path == "/openapi.json" {
			return c.Next()
		}
		const prefix = "Bearer "
		header := c.Get("Authorization")
		if !strings.HasPrefix(header, prefix) {
			return unauthorized(c)
		}
		provided := strings.TrimSpace(strings.TrimPrefix(header, prefix))
		if len(provided) != len(expected) || subtle.ConstantTimeCompare([]byte(provided), []byte(expected)) != 1 {
			return unauthorized(c)
		}
		return c.Next()
	}
}

func unauthorized(c fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": fiber.Map{
			"message": "Missing or invalid API key. Provide Authorization: Bearer YOUR_API_KEY.",
			"type": "authentication_error",
			"code": "invalid_api_key",
		},
	})
}
