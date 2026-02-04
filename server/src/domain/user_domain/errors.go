package user_domain

import "aeroline/src/domain/shared"

var (
	ErrUsernameTooShort = &shared.AppError{
		Type: shared.TypeValidation,
		Msg:  "Username length must be greater than 8",
	}
	ErrUsernameTooLong = &shared.AppError{
		Type: shared.TypeValidation,
		Msg:  "Username length must be less than 32",
	}

	ErrPasswordTooShort = &shared.AppError{
		Type: shared.TypeValidation,
		Msg:  "Password length must be greater than 8",
	}
	ErrPasswordMismath = &shared.AppError{
		Type: shared.TypeUnauthorized,
		Msg:  "Password mismatch",
	}

	ErrUserMissed = &shared.AppError{
		Type: shared.TypeMissingData,
		Msg:  "User data missed in persistense",
	}

	ErrUserWithNameNotFound = &shared.AppError{
		Type: shared.TypeNotFound,
		Msg:  "User with this username not founded",
	}
)
