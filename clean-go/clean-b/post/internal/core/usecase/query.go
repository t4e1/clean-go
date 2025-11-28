package usecase

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/core/domain"
)

// packages for query business logic.
type QueryUsecase struct {
	BaseUC
}

func (q *QueryUsecase) GetPost(postId int) (domain.Post, error) {

	return domain.Post{}, nil
}

func (q *QueryUsecase) GetPosts(query string) ([]domain.Post, error) {

	return nil, nil
}
