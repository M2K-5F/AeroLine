package user_domain

func NewUser(usernameVO Username, password string) (*User, error) {

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
