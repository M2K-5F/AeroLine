package booking

type IReader interface {
	GetTickerByID(id TicketID) (*Ticket, error)
}

type IWriter interface {
	SaveTicket(*Ticket) error
}
