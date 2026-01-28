package user

type User struct {
	id           UserID
	username     Username
	permissions  []Permission
	passwordHash Password
}

func (ths User) ID() UserID {
	return ths.id
}

func (ths User) Username() Username {
	return ths.username
}

func (ths User) Permissions() []Permission {
	return ths.permissions
}
