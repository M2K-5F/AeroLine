package main

import (
	"aeroline/src/infra/persistense/database"
	"aeroline/src/infra/persistense/writers"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/swagger/v2"
)

func initApp() (*fiber.App, func(), error) {
	ctx, cancel := context.WithCancel(context.Background())

	pool, err := database.GetConnectionPoolFromEnv(ctx)
	if err != nil {
		cancel()
		return nil, nil, err
	}

	_ = writers.NewWriter(pool)

	app := fiber.New()
	app.Get("/docs/*", swagger.HandlerDefault)

	api := app.Group("/api")
	api.Get("/health", healthCheck)

	cleanup := func() {
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer shutdownCancel()

		if err := app.ShutdownWithContext(shutdownCtx); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}

		cancel()
	}

	return app, cleanup, nil
}

// @Tags health
// @Success 200 {object} map[string]string "status: ok"
// @Router /health [get]
func healthCheck(c fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}
