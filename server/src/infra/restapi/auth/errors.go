package rest_auth

import "aeroline/src/domain/shared"

var (
	ErrRefreshTokenInvalid = &shared.AppError{
		Type: shared.TypeUnauthorized,
		Msg:  "Invalid refresh token",
	}

	ErrNeedRefresh = &shared.AppError{
		Type: shared.TypeUnauthorized,
		Msg:  "Token outdated",
	}

	ErrTokenNotFound = &shared.AppError{
		Type: shared.TypeUnauthorized,
		Msg:  "Unauthorized",
	}

	ErrInvalidTokenFormat = &shared.AppError{
		Type: shared.TypeUnauthorized,
		Msg:  "Invalid token format",
	}

	ErrNoPermission = &shared.AppError{
		Type: shared.TypeForbidden,
		Msg:  "You not have permission",
	}

	ErrTokenExpired = &shared.AppError{
		Type: shared.TypeUnauthorized,
		Msg:  "Token Expired",
	}

	ErrSessionBlocked = &shared.AppError{
		Type: shared.TypeUnauthorized,
		Msg:  "Your session have been blocked",
	}

	ErrSessionNotFound = &shared.AppError{
		Type: shared.TypeUnauthorized,
		Msg:  "Session not defined",
	}

	ErrUndefinedDevice = &shared.AppError{
		Type: shared.TypeUnauthorized,
		Msg:  "Undefined device. You need relogin",
	}
)
