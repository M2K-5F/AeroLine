package flight_domain

import "aeroline/src/domain/shared"

var (
	ErrAlreadyOccupied = &shared.AppError{
		Msg:  "This Seat is already occupied",
		Type: shared.TypeBusinessLogic,
	}

	ErrFlightMissed = &shared.AppError{
		Type: shared.TypeMissingData,
		Msg:  "Flight data missed in persistense",
	}

	ErrFlightSeatMissed = &shared.AppError{
		Type: shared.TypeMissingData,
		Msg:  "FlightSeat data missed in persistense",
	}
)
