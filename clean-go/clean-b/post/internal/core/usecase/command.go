package usecase

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/core/domain"
	"github.com/t4e1/clean-go/clean-b/post/internal/ports/out"
)

// packages for command business logic.
type CommandUsecase struct {
	out out.PostgresOutPort
}

func NewCommandUsecase(op out.PostgresOutPort) *CommandUsecase {
	return &CommandUsecase{out: op}
}

func (c *CommandUsecase) New(postInfo domain.Post) (bool, error) {

	return false, nil
}

func (c *CommandUsecase) Update(postInfo domain.Post) (bool, error) {

	return false, nil
}

func (c *CommandUsecase) Delete(postId int) (bool, error) {

	return false, nil
}
