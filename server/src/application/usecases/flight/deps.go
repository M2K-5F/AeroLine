package flight_usecases

import (
	"aeroline/src/application/interfaces"
	"aeroline/src/domain/flight"
	"aeroline/src/domain/plane"
)

type DepsModule struct {
	writer    interfaces.IWriter
	planeRdr  plane.IReader
	flightRdr flight.IReader
}

type UseCase struct {
	deps *DepsModule
}

func New(deps *DepsModule) *UseCase {
	return &UseCase{
		deps: deps,
	}
}
