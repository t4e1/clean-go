package usecase

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/core/domain"
	"github.com/t4e1/clean-go/clean-b/post/internal/ports/out"
)

// packages for query business logic.
type QueryUsecase struct {
	Out out.PostgresOutPort
}

func NewQueryUsecase(outAdapter out.PostgresOutPort) QueryUsecase {
	return QueryUsecase{
		Out: outAdapter,
	}
}

func (q *QueryUsecase) GetPost(postId int) (domain.Post, error) {

	return domain.Post{}, nil
}

func (q *QueryUsecase) GetPosts(query string) ([]domain.Post, error) {

	return nil, nil
}
