package user_usecase

import (
	"aeroline/src/application/commands"
	"aeroline/src/domain/user_domain"
	"context"
)

func (ths UseCase) Login(ctx context.Context, cmd commands.LoginCMD) (*user_domain.User, error) {
	username, err := user_domain.NewUsername(cmd.Username)
	if err != nil {
		return nil, err
	}

	usr, err := ths.deps.UserRdr.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if !usr.VerifyPassword(cmd.Password) {
		return nil, user_domain.ErrPasswordMismath
	}

	return usr, nil
}
