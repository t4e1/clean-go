package usecase

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/ports/out"
)

type RESTApi struct {
	QueryUsecase
	CommandUsecase
}

func InitUsecase(outAdapter out.PostgresOutPort) *RESTApi {
	return &RESTApi{
		QueryUsecase:   NewQueryUsecase(outAdapter),
		CommandUsecase: NewCommandUsecase(outAdapter),
	}
}
