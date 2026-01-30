package plane_usecase

import (
	"aeroline/src/domain/plane_domain"
	"aeroline/src/domain/user_domain"
	"context"
)

func (ths UseCase) Create(ctx context.Context, userID user_domain.UserID, name string) (*plane_domain.Plane, error) {
	return nil, nil
}
