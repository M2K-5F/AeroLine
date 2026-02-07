package controllers

import "github.com/gofiber/fiber/v3"

type AppController struct{}

type HealthResponse struct {
	Status string `json:"status" example:"ok"`
}

// @Tags health
// @Success 200 {object} HealthResponse
// @Router /health [get]
func (ths AppController) Health(c fiber.Ctx) error {
	return c.JSON(HealthResponse{Status: "ok"})
}
