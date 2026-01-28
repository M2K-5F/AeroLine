package booking

import "context"

type IReader interface {
	GetTickerByID(ctx context.Context, id TicketID) (*Ticket, error)
}

type IWriter interface {
	SaveTicket(*Ticket) error
}
