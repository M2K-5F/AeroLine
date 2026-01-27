package user

import "aeroline/src/domain/shared"

const (
	ErrUsernameTooShort = shared.DomainError("Username length must be greater than 8")
	ErrUsernameTooLong  = shared.DomainError("Username length must be less than 32")

	ErrPasswordTooShort = shared.DomainError("Password length must be greater than 8")
)
