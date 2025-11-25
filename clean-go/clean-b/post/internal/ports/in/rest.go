package in

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/core/domain"
)

// packages for REST API(HTTP) interface

// interfaces for facade interface set
type InboundRESTApiPort interface {
	QueryPort
	CommandPort
}

// higher complexity => divide port into function level
// ex) PostQueryPort, ApiQueryPort, ....
type QueryPort interface {
	GetPost(postId int) (domain.Post, error)
	GetPosts(query string) ([]domain.Post, error)
}

type CommandPort interface {
	NewPost(postInfo domain.Post) (bool, error)
	UpdatePost(postInfo domain.Post) (bool, error)
	DeletePost(postId int) (bool, error)
}
