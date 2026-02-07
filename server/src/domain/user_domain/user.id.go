package user_domain

import (
	"aeroline/src/domain/shared"
)

type UserID struct {
	shared.ID
}

func NewUserID() UserID {
	return UserID{ID: shared.NewID()}
}
