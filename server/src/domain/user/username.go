package user

type Username string

func (ths Username) String() string {
	return string(ths)
}

func NewUsername(plain string) (Username, error) {
	if len(plain) < 8 {
		return "", ErrUsernameTooShort
	}

	if len(plain) > 32 {
		return "", ErrUsernameTooLong
	}

	return Username(plain), nil
}
