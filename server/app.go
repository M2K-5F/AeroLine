package main

import (
	"aeroline/src/domain/user_domain"
	rest_auth "aeroline/src/infra/restapi/auth"
	"aeroline/src/infra/restapi/handlers"
	"aeroline/src/infra/restapi/middlewares"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/swagger/v2"
)

func initApp() (*fiber.App, func(), error) {
	ctx, cancel := context.WithCancel(context.Background())

	deps, err := ResolveDependencies(ctx)
	if err != nil {
		cancel()
		return nil, nil, err
	}

	app := fiber.New()

	app.Get("/docs/*", swagger.HandlerDefault)
	app.Use(middlewares.Logger())
	app.Use(middlewares.Error())

	api := app.Group("/api")
	api.Use(rest_auth.DeviceIDMiddleware)
	api.Get("/health", handlers.HealthCheck)

	auth := api.Group("/auth")

	auth.Post("/register", deps.AuthController.Register)
	auth.Post("/login", deps.AuthController.Login)
	auth.Get("/me",
		deps.Filter(
			user_domain.AdminPermission,
			user_domain.CustomerPermission,
		),
		deps.AuthController.Me,
	)
	auth.Patch("/refresh", deps.AuthController.Refresh)
	auth.Get("/sessions", deps.Filter(), deps.AuthController.GetSessions)

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
