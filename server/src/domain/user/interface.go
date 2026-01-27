package user

type IReader interface {
	GetUserByID(id UserID) (*User, error)
}

type IWriter interface {
	SaveUser(*User) error
}
