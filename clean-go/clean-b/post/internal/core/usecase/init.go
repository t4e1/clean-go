package usecase

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/ports/out"
)

type Usecases struct {
	Query   *QueryUsecase
	Command *CommandUsecase
}

func InitUsecases(outAdapter out.PostgresOutPort) *Usecases {
	return &Usecases{
		Query:   NewQuerydUsecase(outAdapter),
		Command: NewCommandUsecase(outAdapter),
	}
}
