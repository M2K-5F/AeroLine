package flight

import "context"

type IReader interface {
	GetFlightByID(ctx context.Context, id FlightID) (*Flight, error)
	GetFlightSeatByID(ctx context.Context, id FlightSeatID) (*FlightSeat, error)
}

type IWriter interface {
	SaveFlight(*Flight) error
	SaveFlightSeat(*FlightSeat) error
}
