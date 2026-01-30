package flight_domain

import (
	"aeroline/src/domain/plane_domain"
	"aeroline/src/domain/shared"
)

type FlightSeat struct {
	id         FlightSeatID
	isOccupied bool
	price      shared.Price
	seatID     plane_domain.SeatID
	flightID   FlightID
}

func (ths FlightSeat) ID() FlightSeatID {
	return ths.id
}

func (ths FlightSeat) IsOccupied() bool {
	return ths.isOccupied
}

func (ths FlightSeat) Price() shared.Price {
	return ths.price
}

func (ths FlightSeat) SeatID() plane_domain.SeatID {
	return ths.seatID
}

func (ths FlightSeat) FlightID() FlightID {
	return ths.flightID
}
