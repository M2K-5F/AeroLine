package plane

import "aeroline/src/domain/shared"

const (
	ErrNameTooShort = shared.DomainError("Plane name too short")
	ErrNameTooLong  = shared.DomainError("Plane name too long")

	ErrUnknownSeatClass = shared.DomainError("Unknown seat class")
)
