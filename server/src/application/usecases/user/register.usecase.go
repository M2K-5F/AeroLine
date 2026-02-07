package user_usecase

import (
	"aeroline/src/application/interfaces"
	"aeroline/src/domain/user_domain"
	"context"
)

type RegisterUserCMD struct {
	Username user_domain.Username
	Password string
}

func (ths UseCase) Register(ctx context.Context, cmd RegisterUserCMD) (*user_domain.User, error) {
	user, err := user_domain.NewUser(cmd.Username, cmd.Password)
	if err != nil {
		return nil, err
	}

	if err := ths.deps.Writer.Execute(ctx, func(writer interfaces.ITransactionWriter) error {
		return writer.SaveUser(user)
	}); err != nil {
		return nil, err
	}

	return user, nil
}
