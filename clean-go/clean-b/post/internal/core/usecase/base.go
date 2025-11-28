package usecase

import "github.com/t4e1/clean-go/clean-b/post/internal/ports/out"

type BaseUC struct {
	Out out.PostgresOutPort
}
