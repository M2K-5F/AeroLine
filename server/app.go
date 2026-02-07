package main

import (
	"aeroline/src/domain/user_domain"
	rest_auth "aeroline/src/infra/restapi/auth"
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
	filter := deps.Filter

	app := fiber.New()

	app.Use(middlewares.Logger())
	app.Use(middlewares.Error())
	app.Use(rest_auth.DeviceIDMiddleware)

	app.Get("/docs/*", swagger.HandlerDefault)
	app.Get("health", deps.AppController.Health)

	auth := app.Group("/auth")
	auth.Post("/register", deps.AuthController.Register)
	auth.Post("/login", deps.AuthController.Login)
	auth.Get(`/me`, filter(user_domain.AdminPermission, user_domain.CustomerPermission), deps.AuthController.Me)
	auth.Patch("/refresh", deps.AuthController.Refresh)
	auth.Get("/sessions", filter(), deps.AuthController.GetSessions)
	auth.Post("/logout", filter(), deps.AuthController.Logout)

	users := app.Group("/users")
	users.Get("/:id", filter(), deps.UserController.GetByID)

	flights := app.Group("/flight")
	flights.Get("/cities", deps.FlightController.FindCitiesByName)

	cleanup := func() {
		deps.Close()
		shutdownCtx, shutdownCancel := context.WithTimeout(
			context.Background(),
			30*time.Second,
		)
		defer shutdownCancel()

		if err := app.ShutdownWithContext(shutdownCtx); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}

		cancel()
	}

	return app, cleanup, nil
}
