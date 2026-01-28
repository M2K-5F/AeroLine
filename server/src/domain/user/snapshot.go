package user

type Snapshot struct {
	ID           UserID
	Username     Username
	PasswordHash Password
	Permissions  []Permission
}

func (ths User) Snapshot() Snapshot {
	return Snapshot{
		ID:           ths.id,
		Username:     ths.username,
		Permissions:  ths.permissions,
		PasswordHash: ths.passwordHash,
	}
}

func Restore(spt Snapshot) *User {
	return &User{
		id:          spt.ID,
		username:    spt.Username,
		permissions: spt.Permissions,
	}
}
