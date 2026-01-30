package models

import (
	"aeroline/src/domain/booking_domain"
	"aeroline/src/domain/flight_domain"
	"aeroline/src/domain/plane_domain"
	"aeroline/src/domain/user_domain"
	"time"
)

type UserRow struct {
	ID           user_domain.UserID `db:"id"`
	Username     string             `db:"username"`
	Permissions  []string           `db:"permissions"`
	PasswordHash string             `db:"password"`
}

type PlaneRow struct {
	ID   plane_domain.PlaneID `db:"id"`
	Name string               `db:"name"`
}

type SeatRow struct {
	ID      plane_domain.SeatID  `db:"id"`
	PlaneID plane_domain.PlaneID `db:"plane_id"`
	Tag     string               `db:"tag"`
	Serial  int                  `db:"serial"`
	Class   string               `db:"class"`
}

type FlightRow struct {
	ID            flight_domain.FlightID `db:"id"`
	DepartureCode string                 `db:"departure_aip_code"`
	ArrivalCode   string                 `db:"arrival_aip_code"`
	PlaneID       plane_domain.PlaneID   `db:"plane_id"`
	ArrivalTime   time.Time              `db:"arrival_time"`
	DepartureTime time.Time              `db:"departure_time"`
}

type FlightSeatRow struct {
	ID         flight_domain.FlightSeatID `db:"id"`
	IsOccupied bool                       `db:"is_occupied"`
	PriceRow
	SeatID   plane_domain.SeatID    `db:"seat_id"`
	FlightID flight_domain.FlightID `db:"flight_id"`
}

type TicketRow struct {
	ID      booking_domain.TicketID `db:"id"`
	BuyerID user_domain.UserID      `db:"buyer_id"`
	PriceRow
	FlightSeatID flight_domain.FlightSeatID `db:"flight_seat_id"`
}

type PriceRow struct {
	Amount   int64  `db:"amount"`
	Currency string `db:"currency"`
}
