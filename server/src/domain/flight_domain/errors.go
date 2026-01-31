package flight_domain

import "aeroline/src/domain/shared"

var (
	ErrAlreadyOccupied = &shared.AppError{
		Msg:  "This Seat is already occupied",
		Type: shared.TypeBusinessLogic,
	}
)
