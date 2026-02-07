package rest_auth

import (
	"aeroline/src/domain/user_domain"
	"slices"

	"github.com/gofiber/fiber/v3"
)

type AuthUser struct {
	userID      user_domain.UserID
	permissions []user_domain.Permission
}

type PermissionFilter func(permissions ...user_domain.Permission) fiber.Handler

func NewPermissionFilter(service AuthService) PermissionFilter {
	return func(permissions ...user_domain.Permission) fiber.Handler {
		return func(c fiber.Ctx) error {
			authUser, err := getAuthUserFromRequest(c, service)
			if err != nil {
				return err
			}

			if slices.ContainsFunc(authUser.permissions, func(p user_domain.Permission) bool {
				return len(permissions) == 0 || slices.Contains(permissions, p)
			}) {
				c.Locals("user-id", authUser.userID)
				return c.Next()
			}

			return ErrNoPermission
		}
	}
}

func getAuthUserFromRequest(c fiber.Ctx, service AuthService) (*AuthUser, error) {
	authUser, ok := c.Locals("auth-user").(AuthUser)
	if !ok {
		token, err := getBearerFromInLocals(c)
		if err != nil {
			return nil, err
		}
		userID, permissions, err := service.VerifyAccessToken(c.Context(), AccessToken(token))
		if err != nil {
			return nil, err
		}

		authUser = AuthUser{
			permissions: *permissions,
			userID:      *userID,
		}
	}

	return &authUser, nil
}

func getBearerFromInLocals(c fiber.Ctx) (string, error) {
	authHeaders, ok := c.GetReqHeaders()["Authorization"]
	if !ok || len(authHeaders) == 0 {
		return "", ErrTokenNotFound
	}

	authHeader := authHeaders[0]

	// words := strings.Split(strings.TrimSpace(authHeader), " ")
	// if len(words) != 2 || strings.ToLower(words[0]) != "bearer" {
	// 	return "", ErrInvalidTokenFormat
	// }

	// return words[1], nil
	return authHeader, nil
}
