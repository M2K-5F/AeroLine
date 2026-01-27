package flight

type IReader interface {
	GetFlightByID(id FlightID) (*Flight, error)
	GetFlightSeatByID(id FlightSeatID) (*FlightSeat, error)
}

type IWriter interface {
	SaveFlight(*Flight) error
	SaveFlightSeat(*FlightSeat) error
}
