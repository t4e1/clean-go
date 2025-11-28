package usecase

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/core/domain"
)

// packages for command business logic.
type CommandUsecase struct {
	BaseUC
}

func (c *CommandUsecase) NewPost(postInfo domain.Post) (bool, error) {

	return false, nil
}

func (c *CommandUsecase) UpdatePost(postInfo domain.Post) (bool, error) {

	return false, nil
}

func (c *CommandUsecase) DeletePost(postId int) (bool, error) {

	return false, nil
}
