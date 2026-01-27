package flight

import (
	"aeroline/src/domain/plane"
	"aeroline/src/domain/shared"
	"time"
)

type NewFlightData struct {
	Plane              *plane.Plane
	Seats              []*plane.Seat
	FlightTag          string
	ArrivalTime        time.Time
	DepartureTime      time.Time
	ArrivalAirport     *shared.Airport
	DepartureAirport   *shared.Airport
	BusinessClassPrice shared.Price
	EconomyClassPrice  shared.Price
}

func NewFlight(data *NewFlightData) (*Flight, []*FlightSeat, error) {
	flight, err := newFlight(data)
	if err != nil {
		return nil, nil, err
	}

	seats := make([]*FlightSeat, len(data.Seats))

	for i, seat := range data.Seats {
		var flightSeat *FlightSeat
		var err error

		switch seat.Class() {
		case plane.BusinessSeat:
			flightSeat, err = newFlightSeat(
				flight, seat, data.BusinessClassPrice,
			)

		case plane.EconomySeat:
			flightSeat, err = newFlightSeat(
				flight, seat, data.EconomyClassPrice,
			)

		default:
			return nil, nil, plane.ErrUnknownSeatClass
		}

		if err != nil {
			return nil, nil, err
		}
		seats[i] = flightSeat
	}

	return flight, seats, nil
}

func newFlight(data *NewFlightData) (*Flight, error) {
	return &Flight{
		id:            FlightID{ID: shared.NewID()},
		planeID:       data.Plane.ID(),
		departure:     data.DepartureAirport,
		arrival:       data.ArrivalAirport,
		departureTime: data.DepartureTime,
		arrivalTime:   data.ArrivalTime,
	}, nil
}

func newFlightSeat(flight *Flight, seat *plane.Seat, price shared.Price) (*FlightSeat, error) {
	return &FlightSeat{
		id:         FlightSeatID{ID: shared.NewID()},
		flightID:   flight.ID(),
		isOccupied: false,
		seatID:     seat.ID(),
		price:      price,
	}, nil
}
