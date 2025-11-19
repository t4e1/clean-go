package usecases

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/core/domain"
)

// packages for business logic.
type RESTApi struct {
}

func (*RESTApi) GetPost(postId int) (domain.Post, error) {

	return domain.Post{}, nil
}

func (*RESTApi) GetPosts(query string) ([]domain.Post, error) {

	return nil, nil
}

func (*RESTApi) NewPost(postInfo domain.Post) (bool, error) {

	return false, nil
}

func (*RESTApi) UpdatePost(postInfo domain.Post) (bool, error) {

	return false, nil
}

func (*RESTApi) DeletePost(postId int) (bool, error) {

	return false, nil
}
