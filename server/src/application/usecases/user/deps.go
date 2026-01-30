package user_usecase

import (
	"aeroline/src/application/interfaces"
	"aeroline/src/domain/user_domain"
)

type DepsModule struct {
	Writer  interfaces.IWriter
	UserRdr user_domain.IReader
}

type UseCase struct {
	deps *DepsModule
}

func New(deps *DepsModule) *UseCase {
	return &UseCase{
		deps: deps,
	}
}
