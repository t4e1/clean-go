package usecase

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/core/domain"
	"github.com/t4e1/clean-go/clean-b/post/internal/ports/out"
)

// packages for query business logic.
type QueryUsecase struct {
	out out.PostgresOutPort
}

func NewQuerydUsecase(op out.PostgresOutPort) *QueryUsecase {
	return &QueryUsecase{out: op}
}

func (q *QueryUsecase) Get(postId int) (domain.Post, error) {

	return domain.Post{}, nil
}

func (q *QueryUsecase) GetLists(query string) ([]domain.Post, error) {

	return nil, nil
}
