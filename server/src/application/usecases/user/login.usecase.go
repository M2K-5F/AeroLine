package user_usecase

import (
	"aeroline/src/application/commands"
	"aeroline/src/domain/shared"
	"aeroline/src/domain/user_domain"
	"context"
	"errors"
)

func (ths UseCase) Login(ctx context.Context, cmd commands.LoginCMD) (*user_domain.User, error) {
	username, err := user_domain.NewUsername(cmd.Username)
	if err != nil {
		return nil, err
	}

	usr, err := ths.deps.UserRdr.GetUserByUsername(ctx, username)
	if errors.Is(err, shared.ErrDataNotFound) {
		return nil, user_domain.ErrUserWithNameNotFound
	}

	if err != nil {
		return nil, err
	}

	if !usr.VerifyPassword(cmd.Password) {
		return nil, user_domain.ErrPasswordMismath
	}

	return usr, nil
}
