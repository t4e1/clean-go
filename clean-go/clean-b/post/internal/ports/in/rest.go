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
	Get(postId int) (domain.Post, error)
	GetLists(query string) ([]domain.Post, error)
}

type CommandPort interface {
	New(postInfo domain.Post) (bool, error)
	Update(postInfo domain.Post) (bool, error)
	Delete(postId int) (bool, error)
}
