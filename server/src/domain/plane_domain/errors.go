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
)
