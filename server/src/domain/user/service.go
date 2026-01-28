package user

func NewUser(username string, password string) (*User, error) {
	usernameVO, err := NewUsername(username)
	if err != nil {
		return nil, err
	}

	passwordVO, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		id:           NewUserID(),
		username:     usernameVO,
		permissions:  []Permission{CustomerPermission},
		passwordHash: passwordVO,
	}, nil
}
