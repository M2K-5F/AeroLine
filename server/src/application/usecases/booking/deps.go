package booking_usecase

import (
	"aeroline/src/application/interfaces"
	"aeroline/src/domain/booking_domain"
	"aeroline/src/domain/flight_domain"
	"aeroline/src/domain/user_domain"
)

type DepsModule struct {
	writer     interfaces.IWriter
	userRdr    user_domain.IReader
	flightRdr  flight_domain.IReader
	bookingRdr booking_domain.IReader
}

type UseCase struct {
	deps *DepsModule
}

func New(deps *DepsModule) *UseCase {
	return &UseCase{
		deps: deps,
	}
}
