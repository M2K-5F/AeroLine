package booking_usecase

import (
	"aeroline/src/application/interfaces"
	"aeroline/src/domain/booking"
	"aeroline/src/domain/flight"
	"aeroline/src/domain/user"
)

type DepsModule struct {
	writer     interfaces.IWriter
	userRdr    user.IReader
	flightRdr  flight.IReader
	bookingRdr booking.IReader
}

type UseCase struct {
	deps *DepsModule
}

func New(deps *DepsModule) *UseCase {
	return &UseCase{
		deps: deps,
	}
}
