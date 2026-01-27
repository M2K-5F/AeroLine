package plane_usecase

import (
	"aeroline/src/application/interfaces"
	"aeroline/src/domain/plane"
	"aeroline/src/domain/user"
)

type DepsModule struct {
	writer   interfaces.IWriter
	userRdr  user.IReader
	planeRdr plane.IReader
}

type UseCase struct {
	deps *DepsModule
}

func New(deps *DepsModule) *UseCase {
	return &UseCase{
		deps: deps,
	}
}
