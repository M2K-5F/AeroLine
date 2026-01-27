package user

func NewUser(username string) (*User, error) {
	usernameVO, err := NewUsername(username)
	if err != nil {
		return nil, err
	}

	return &User{
		id:          NewUserID(),
		username:    usernameVO,
		permissions: []Permission{CustomerPermission},
	}, nil
}

func Restore(spt Snapshot) *User {
	permissions := make([]Permission, len(spt.Permissions))

	for i, permission := range spt.Permissions {
		permissions[i] = Permission(permission)
	}

	return &User{
		id:          spt.ID,
		username:    Username(spt.Username),
		permissions: permissions,
	}
}
