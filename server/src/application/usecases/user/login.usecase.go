package user_usecase

import (
	"aeroline/src/domain/shared"
	domain "aeroline/src/domain/user_domain"
	"context"
	"errors"
)

type LoginCMD struct {
	Username domain.Username
	Password string
}

func (ths UseCase) Login(ctx context.Context, cmd LoginCMD) (*domain.User, error) {
	usr, err := ths.deps.UserRdr.GetUserByUsername(ctx, cmd.Username)
	if errors.Is(err, shared.ErrDataNotFound) {
		return nil, domain.ErrUserWithNameNotFound
	}

	if err != nil {
		return nil, err
	}

	if !usr.VerifyPassword(cmd.Password) {
		return nil, domain.ErrPasswordMismath
	}

	return usr, nil
}
