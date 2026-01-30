package plane_domain

type Name string

func (ths Name) String() string {
	return string(ths)
}

func NewName(plain string) (Name, error) {
	if len(plain) < 8 {
		return "", ErrNameTooShort
	}

	if len(plain) > 64 {
		return "", ErrNameTooLong
	}

	return Name(plain), nil
}
