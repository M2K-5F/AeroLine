package handlers

import (
	"github.com/gofiber/fiber/v3"
)

// @Tags health
// @Success 200 {object} map[string]string "status: ok"
// @Router /health [get]
func HealthCheck(c fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}
