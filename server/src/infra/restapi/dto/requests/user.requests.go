package requests

import (
	uc "aeroline/src/application/usecases/user"
	domain "aeroline/src/domain/user_domain"
)

type RegisterUserRequest struct {
	Username string
	Password string
}

func (ths RegisterUserRequest) ToCMD() (*uc.RegisterUserCMD, error) {
	username, err := domain.NewUsername(ths.Username)
	if err != nil {
		return nil, err
	}

	return &uc.RegisterUserCMD{
		Username: username,
		Password: ths.Password,
	}, nil
}

type LoginUserRequest struct {
	Username string
	Password string
}

func (ths LoginUserRequest) ToCMD() (*uc.LoginCMD, error) {
	username, err := domain.NewUsername(ths.Username)
	if err != nil {
		return nil, err
	}

	return &uc.LoginCMD{
		Username: username,
		Password: ths.Password,
	}, nil
}

type GetUserByIDRequest struct {
	UserID string `uri:"id"`
}

func (ths GetUserByIDRequest) ToCMD() (*uc.GetByIdCMD, error) {
	userID := new(domain.UserID)
	if err := userID.Parse(ths.UserID); err != nil {
		return nil, err
	}

	return &uc.GetByIdCMD{
		UserID: *userID,
	}, nil
}
