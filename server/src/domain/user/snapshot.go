package user

type Snapshot struct {
	ID          UserID
	Username    string
	Permissions []string
}

func (ths User) Snapshot() Snapshot {
	permissions := make([]string, len(ths.permissions))

	for i, permission := range ths.permissions {
		permissions[i] = permission.String()
	}

	return Snapshot{
		ID:          ths.id,
		Username:    ths.username.String(),
		Permissions: permissions,
	}
}
