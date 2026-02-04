package main

import (
	user_usecase "aeroline/src/application/usecases/user"
	"aeroline/src/infra/persistense/database"
	"aeroline/src/infra/persistense/readers"
	"aeroline/src/infra/persistense/writers"
	rest_auth "aeroline/src/infra/restapi/auth"
	"aeroline/src/infra/restapi/controllers"
	"context"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v3"
)

type Filters struct {
	AdminFilter    fiber.Handler
	CustomerFilter fiber.Handler
}

type Dependencies struct {
	AuthController *controllers.AuthController
	Filter         rest_auth.PermissionFilter
	Close          func()
}

func ResolveDependencies(ctx context.Context) (*Dependencies, error) {
	pool, err := database.GetConnectionPoolFromEnv(ctx)
	if err != nil {
		return nil, err
	}

	dsn := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	client := redis.NewClient(&redis.Options{
		Addr: dsn,
		DB:   0,
	})

	writer := writers.NewWriter(pool)

	userReader := readers.NewUserReader(pool)
	// planeReader := readers.NewPlaneReader(pool)
	// flightReader := readers.NewFlightReader(pool)
	// bookingReader := readers.NewBookingReader(pool)

	tokenService := rest_auth.NewAuthService(client, rest_auth.NewConfigFromEnv(), userReader)

	filter := rest_auth.NewPermissionFilter(*tokenService)

	userUC := user_usecase.New(&user_usecase.DepsModule{
		Writer:  writer,
		UserRdr: userReader,
	})

	authCtrlr := controllers.NewAuthController(userUC, tokenService)

	return &Dependencies{
		AuthController: authCtrlr,
		Filter:         filter,
		Close: func() {
			pool.Close()
			client.Close()
		},
	}, nil
}
