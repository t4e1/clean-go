package out

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/core/domain"
)

// packages for database interface
type PostgresOutPort interface {
	SelectPost() (domain.Post, error)
}
