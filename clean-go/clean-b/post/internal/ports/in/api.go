package in

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/core/domain"
)

// packages for REST API(HTTP) interface

type InboundRESTApiPort interface {
	GetPost(postId int) (domain.Post, error)
	GetPosts(query string) ([]domain.Post, error)
	NewPost(postInfo domain.Post) (bool, error)
	UpdatePost(postInfo domain.Post) (bool, error)
	DeletePost(postId int) (bool, error)
}
