package flight

import (
	"aeroline/src/domain/plane"
	"aeroline/src/domain/shared"
	"time"
)

type Flight struct {
	id            FlightID
	departure     *shared.Airport
	arrival       *shared.Airport
	planeID       plane.PlaneID
	arrivalTime   time.Time
	departureTime time.Time
}

func (ths Flight) ID() FlightID {
	return ths.id
}

func (ths Flight) Departure() shared.Airport {
	return *ths.departure
}

func (ths Flight) Arrival() shared.Airport {
	return *ths.arrival
}

func (ths Flight) DepartureTime() time.Time {
	return ths.departureTime
}

func (ths Flight) ArrivalTime() time.Time {
	return ths.arrivalTime
}
