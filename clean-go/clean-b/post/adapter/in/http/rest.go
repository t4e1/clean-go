package in

import "github.com/t4e1/clean-go/clean-b/post/internal/core/usecase"

// packages for REST API implementation

type RESTInAdapter struct {
	usecase usecase.RESTApi
}

type HttpHandler interface {
}

func InitRESTAdapter(uc usecase.RESTApi) *RESTInAdapter {

	return &RESTInAdapter{
		usecase: uc,
	}
}
