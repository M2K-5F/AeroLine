package booking

import (
	"aeroline/src/domain/flight"
	"aeroline/src/domain/shared"
	"aeroline/src/domain/user"
)

type Ticket struct {
	id           TicketID
	buyerID      user.UserID
	price        shared.Price
	flightSeatID flight.FlightSeatID
}

func (ths Ticket) ID() TicketID {
	return ths.id
}
