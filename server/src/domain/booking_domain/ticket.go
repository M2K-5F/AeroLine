package booking_domain

import (
	"aeroline/src/domain/flight_domain"
	"aeroline/src/domain/shared"
	"aeroline/src/domain/user_domain"
)

type Ticket struct {
	id           TicketID
	buyerID      user_domain.UserID
	price        shared.Price
	flightSeatID flight_domain.FlightSeatID
}

func (ths Ticket) ID() TicketID {
	return ths.id
}
