package usecase

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/ports/out"
)

type Usecases struct {
	QueryUsecase
	CommandUsecase
}

func InitUsecases(outAdapter out.PostgresOutPort) *Usecases {
	return &Usecases{
		QueryUsecase:   *NewQuerydUsecase(outAdapter),
		CommandUsecase: *NewCommandUsecase(outAdapter),
	}
}
