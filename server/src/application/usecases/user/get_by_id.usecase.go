package user_usecase

import (
	"aeroline/src/domain/user_domain"
	"context"
)

type GetByIdCMD struct {
	UserID user_domain.UserID
}

func (ths UseCase) GetByID(ctx context.Context, cmd GetByIdCMD) (*user_domain.User, error) {
	user, err := ths.deps.UserRdr.GetUserByID(ctx, cmd.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
