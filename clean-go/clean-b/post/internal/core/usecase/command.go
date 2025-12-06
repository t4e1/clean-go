package usecase

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/core/domain"
	"github.com/t4e1/clean-go/clean-b/post/internal/ports/out"
)

// packages for command business logic.
type CommandUsecase struct {
	Out out.PostgresOutPort
}

func NewCommandUsecase(outAdapter out.PostgresOutPort) CommandUsecase {
	return CommandUsecase{
		Out: outAdapter,
	}
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
