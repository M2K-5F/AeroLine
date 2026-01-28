package models

import (
	"aeroline/src/domain/booking"
	"aeroline/src/domain/flight"
	"aeroline/src/domain/plane"
	"aeroline/src/domain/user"
	"time"
)

type UserRow struct {
	ID           user.UserID `db:"id"`
	Username     string      `db:"username"`
	Permissions  []string    `db:"permissions"`
	PasswordHash string      `db:"password"`
}

type PlaneRow struct {
	ID   plane.PlaneID `db:"id"`
	Name string        `db:"name"`
}

type SeatRow struct {
	ID      plane.SeatID  `db:"id"`
	PlaneID plane.PlaneID `db:"plane_id"`
	Tag     string        `db:"tag"`
	Serial  int           `db:"serial"`
	Class   string        `db:"class"`
}

type FlightRow struct {
	ID            flight.FlightID `db:"id"`
	DepartureCode string          `db:"departure_aip_code"`
	ArrivalCode   string          `db:"arrival_aip_code"`
	PlaneID       plane.PlaneID   `db:"plane_id"`
	ArrivalTime   time.Time       `db:"arrival_time"`
	DepartureTime time.Time       `db:"departure_time"`
}

type FlightSeatRow struct {
	ID         flight.FlightSeatID `db:"id"`
	IsOccupied bool                `db:"is_occupied"`
	PriceRow
	SeatID   plane.SeatID    `db:"seat_id"`
	FlightID flight.FlightID `db:"flight_id"`
}

type TicketRow struct {
	ID      booking.TicketID `db:"id"`
	BuyerID user.UserID      `db:"buyer_id"`
	PriceRow
	FlightSeatID flight.FlightSeatID `db:"flight_seat_id"`
}

type PriceRow struct {
	Amount   int64  `db:"amount"`
	Currency string `db:"currency"`
}
