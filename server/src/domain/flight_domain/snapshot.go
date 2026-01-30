package flight_domain

import (
	"aeroline/src/domain/plane_domain"
	"aeroline/src/domain/shared"
	"time"
)

type FlightSnapshot struct {
	ID            FlightID
	Departure     *shared.Airport
	Arrival       *shared.Airport
	PlaneID       plane_domain.PlaneID
	ArrivalTime   time.Time
	DepartureTime time.Time
}

func (ths Flight) Snapshot() FlightSnapshot {
	return FlightSnapshot{
		ID:            ths.id,
		PlaneID:       ths.planeID,
		Departure:     ths.departure,
		Arrival:       ths.arrival,
		ArrivalTime:   ths.arrivalTime,
		DepartureTime: ths.departureTime,
	}
}

func RestoreFlight(spt FlightSnapshot) *Flight {
	return &Flight{
		id:            spt.ID,
		planeID:       spt.PlaneID,
		arrival:       spt.Arrival,
		departure:     spt.Departure,
		arrivalTime:   spt.ArrivalTime,
		departureTime: spt.DepartureTime,
	}
}

type FlightSeatSnapshot struct {
	ID         FlightSeatID
	IsOccupied bool
	Price      shared.Price
	SeatID     plane_domain.SeatID
	FlightID   FlightID
}

func (ths FlightSeat) Snapshot() FlightSeatSnapshot {
	return FlightSeatSnapshot{
		ID:         ths.id,
		IsOccupied: ths.isOccupied,
		Price:      ths.price,
		SeatID:     ths.seatID,
		FlightID:   ths.flightID,
	}
}

func RestoreFlightSeat(spt FlightSeatSnapshot) *FlightSeat {
	return &FlightSeat{
		id:         spt.ID,
		flightID:   spt.FlightID,
		seatID:     spt.SeatID,
		price:      spt.Price,
		isOccupied: spt.IsOccupied,
	}
}
