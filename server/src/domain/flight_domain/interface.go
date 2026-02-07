package flight_domain

import "context"

type IReader interface {
	GetFlightByID(ctx context.Context, id FlightID) (*Flight, error)
	GetFlightSeatByID(ctx context.Context, id FlightSeatID) (*FlightSeat, error)
}

type IWriter interface {
	SaveFlight(*Flight, ...*Flight) error
	SaveFlightSeat(*FlightSeat, ...*FlightSeat) error
}
