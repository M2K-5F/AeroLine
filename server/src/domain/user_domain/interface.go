package user_domain

import "context"

type IReader interface {
	GetUserByID(ctx context.Context, id UserID) (*User, error)
	GetUserByUsername(ctx context.Context, username Username) (*User, error)
}

type IWriter interface {
	SaveUser(*User) error
}
