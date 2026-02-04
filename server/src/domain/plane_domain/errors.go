package plane_domain

import (
	"aeroline/src/domain/shared"
)

var (
	ErrNameTooShort = &shared.AppError{
		Type: shared.TypeValidation,
		Msg:  "Plane name too short",
	}
	ErrNameTooLong = &shared.AppError{
		Type: shared.TypeValidation,
		Msg:  "Plane name too long",
	}

	ErrUnknownSeatClass = &shared.AppError{
		Type: shared.TypeValidation,
		Msg:  "Unknown seat class",
	}

	ErrPlaneMissed = &shared.AppError{
		Type: shared.TypeMissingData,
		Msg:  "Plane data missed in persistense",
	}

	ErrSeatMissed = &shared.AppError{
		Type: shared.TypeMissingData,
		Msg:  "Seat data missed in persistense",
	}
)
