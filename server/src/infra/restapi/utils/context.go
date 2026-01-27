package rest_utils

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
)

func BaseRequestContext(c fiber.Ctx) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Context(), time.Second)
	return ctx, cancel
}

func CustomRequestContext(c fiber.Ctx, timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Context(), timeout)
	return ctx, cancel
}
