package fligth

import (
	"aeroline/src/domain/plane"
	"aeroline/src/domain/shared"
)

type FlightSeat struct {
	id         FlightSeatID
	isOccupied bool
	price      shared.Price
	seatID     plane.SeatID
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

func (ths FlightSeat) SeatID() plane.SeatID {
	return ths.seatID
}

func (ths FlightSeat) FlightID() FlightID {
	return ths.flightID
}
