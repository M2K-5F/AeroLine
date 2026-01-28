package booking

import (
	"aeroline/src/domain/flight"
	"aeroline/src/domain/shared"
	"aeroline/src/domain/user"
)

type TicketSnapshot struct {
	ID           TicketID
	BuyerID      user.UserID
	Price        shared.Price
	FlightSeatID flight.FlightSeatID
}

func (ths Ticket) Snapshot() TicketSnapshot {
	return TicketSnapshot{
		ID:           ths.id,
		BuyerID:      ths.buyerID,
		Price:        ths.price,
		FlightSeatID: ths.flightSeatID,
	}
}

func RestoreTicket(spt TicketSnapshot) *Ticket {
	return &Ticket{
		id:           spt.ID,
		buyerID:      spt.BuyerID,
		price:        spt.Price,
		flightSeatID: spt.FlightSeatID,
	}
}
