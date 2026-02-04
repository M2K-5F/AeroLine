package booking_domain

import "aeroline/src/domain/shared"

var (
	ErrTicketMissed = &shared.AppError{
		Type: shared.TypeMissingData,
		Msg:  "Ticket data missed in persistense",
	}
)
