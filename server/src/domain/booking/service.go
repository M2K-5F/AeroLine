package booking

import (
	"aeroline/src/domain/flight"
	"aeroline/src/domain/shared"
	"aeroline/src/domain/user"
)

func NewTicket(flightSeat *flight.FlightSeat, user *user.User) (*Ticket, error) {
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
