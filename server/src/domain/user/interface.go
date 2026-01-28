package user

import "context"

type IReader interface {
	GetUserByID(ctx context.Context, id UserID) (*User, error)
}

type IWriter interface {
	SaveUser(*User) error
}
