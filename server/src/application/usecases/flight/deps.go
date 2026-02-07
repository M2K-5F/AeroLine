package flight_usecases

import (
	"aeroline/src/application/interfaces"
	"aeroline/src/domain/flight_domain"
	"aeroline/src/domain/plane_domain"
)

type DepsModule struct {
	Writer    interfaces.IWriter
	PlaneRdr  plane_domain.IReader
	FlightRdr flight_domain.IReader
}

type UseCase struct {
	deps *DepsModule
}

func New(deps *DepsModule) *UseCase {
	return &UseCase{
		deps: deps,
	}
}
