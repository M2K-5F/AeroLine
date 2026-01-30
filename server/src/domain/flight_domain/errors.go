package flight_domain

import "aeroline/src/domain/shared"

const (
	ErrAlreadyOccupied = shared.DomainError("This Seat is already occupied")
)
