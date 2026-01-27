package user

import (
	"aeroline/src/domain/shared"

	"github.com/google/uuid"
)

type UserID uuid.UUID

func (ths UserID) String() string {
	return uuid.UUID(ths).String()
}

func (ths *UserID) Scan(value any) error {
	if value == nil {
		return shared.ErrIDCorrupted
	}

	var id uuid.UUID
	if err := id.Scan(value); err != nil {
		return shared.ErrIDCorrupted
	}

	*ths = UserID(id)
	return nil
}

func NewUserID() UserID {
	return UserID(uuid.New())
}
