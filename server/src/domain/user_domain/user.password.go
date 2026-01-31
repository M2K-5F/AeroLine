package user_domain

import "golang.org/x/crypto/bcrypt"

type Password string

func NewPassword(plain string) (Password, error) {
	if len(plain) < 8 {
		return "", ErrPasswordTooShort
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return Password(bytes), err
}

func (ths Password) Verify(plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(ths), []byte(plain)) == nil
}
