package shared

import (
	"errors"

	"github.com/google/uuid"
)

type ID uuid.UUID

func (ths ID) String() string {
	return uuid.UUID(ths).String()
}

func (ths *ID) Scan(value any) error {
	if value == nil {
		return ErrIDCorrupted
	}

	var id uuid.UUID
	if err := id.Scan(value); err != nil {
		return ErrIDCorrupted
	}

	*ths = ID(id)
	return nil
}

func (ths *ID) Parse(value any) error {
	if value == nil {
		return ErrValidateID
	}

	var id uuid.UUID
	if err := id.Scan(value); err != nil {
		return ErrValidateID
	}

	*ths = ID(id)
	return nil
}

func NewID() ID {
	return ID(uuid.New())
}

var (
	ErrIDCorrupted = errors.New("INTEGRITY ALERT: corrupted ID in DB")
	ErrValidateID  = &AppError{
		Type: TypeValidation,
		Msg:  "Id is not valid",
	}
)
