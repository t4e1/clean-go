package usecase

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/core/domain"
	out "github.com/t4e1/clean-go/clean-b/post/internal/ports/out/rdbms"
)

// packages for business logic.
type RESTApi struct {
	postgresOut rdbms.PostgresOutPort
}

func InitRESTUsecase(outAdapter out.PostgresOutPort) *RESTApi {
	return &RESTApi{
		postgresOut: outAdapter,
	}
}

func (r *RESTApi) GetPost(postId int) (domain.Post, error) {

	return domain.Post{}, nil
}

func (r *RESTApi) GetPosts(query string) ([]domain.Post, error) {

	return nil, nil
}

func (r *RESTApi) NewPost(postInfo domain.Post) (bool, error) {

	return false, nil
}

func (r *RESTApi) UpdatePost(postInfo domain.Post) (bool, error) {

	return false, nil
}

func (r *RESTApi) DeletePost(postId int) (bool, error) {

	return false, nil
}
