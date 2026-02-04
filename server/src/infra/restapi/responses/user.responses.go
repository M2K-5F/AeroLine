package responses

import (
	"aeroline/src/domain/shared"
	"aeroline/src/domain/user_domain"
)

type UserResponse struct {
	ID          string
	Username    string
	Permissions []string
}

func UserToResponse(usr user_domain.User) UserResponse {
	return UserResponse{
		ID:       usr.ID().String(),
		Username: usr.Username().String(),
		Permissions: shared.Map(usr.Permissions(),
			func(p user_domain.Permission) string {
				return p.String()
			},
		),
	}
}

type LoginUserResponse struct {
	UserResponse
	Access string
}
