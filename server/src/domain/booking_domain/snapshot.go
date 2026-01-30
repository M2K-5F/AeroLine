package booking_domain

import (
	"aeroline/src/domain/flight_domain"
	"aeroline/src/domain/shared"
	"aeroline/src/domain/user_domain"
)

type TicketSnapshot struct {
	ID           TicketID
	BuyerID      user_domain.UserID
	Price        shared.Price
	FlightSeatID flight_domain.FlightSeatID
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
