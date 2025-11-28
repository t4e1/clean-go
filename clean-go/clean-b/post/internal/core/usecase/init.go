package usecase

import (
	"github.com/t4e1/clean-go/clean-b/post/internal/core/usecase"
	"github.com/t4e1/clean-go/clean-b/post/internal/ports/out"
)

type RESTApi struct {
	outPort out.PostgresOutPort
	QueryUsecase
	CommandUsecase
}

func InitUsecase(outAdapter out.PostgresOutPort) *RESTApi {
	return &RESTApi{
		outPort:        outAdapter,
		QueryUsecase:   usecase.QueryUsecase,
		CommandUsecase: usecase.CommandUsecase,
	}
}
