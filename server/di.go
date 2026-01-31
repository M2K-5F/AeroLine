package main

import (
	user_usecase "aeroline/src/application/usecases/user"
	"aeroline/src/infra/persistense/database"
	"aeroline/src/infra/persistense/readers"
	"aeroline/src/infra/persistense/writers"
	"aeroline/src/infra/restapi/controllers"
	"context"
)

type Dependencies struct {
	AuthController *controllers.AuthController
}

func ResolveDependencies(ctx context.Context) (*Dependencies, error) {
	pool, err := database.GetConnectionPoolFromEnv(ctx)
	if err != nil {
		return nil, err
	}

	writer := writers.NewWriter(pool)

	userReader := readers.NewUserReader(pool)
	// planeReader := readers.NewPlaneReader(pool)
	// flightReader := readers.NewFlightReader(pool)
	// bookingReader := readers.NewBookingReader(pool)

	userUC := user_usecase.New(&user_usecase.DepsModule{
		Writer:  writer,
		UserRdr: userReader,
	})

	authCtrlr := controllers.NewAuthController(userUC)

	return &Dependencies{
		AuthController: authCtrlr,
	}, nil
}
