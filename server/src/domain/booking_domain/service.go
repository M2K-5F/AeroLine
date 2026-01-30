package booking_domain

import (
	"aeroline/src/domain/flight_domain"
	"aeroline/src/domain/shared"
	"aeroline/src/domain/user_domain"
)

func NewTicket(flightSeat *flight_domain.FlightSeat, user *user_domain.User) (*Ticket, error) {
	if err := flightSeat.OccupieSeat(); err != nil {
		return nil, err
	}

	ticket := Ticket{
		id:           TicketID{ID: shared.NewID()},
		buyerID:      user.ID(),
		flightSeatID: flightSeat.ID(),
		price:        flightSeat.Price(),
	}
	return &ticket, nil
}
