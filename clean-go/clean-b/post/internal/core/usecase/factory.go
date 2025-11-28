package usecase

import "github.com/t4e1/clean-go/clean-b/post/internal/ports/out"

type Factory struct {
}

func NewFactory(out out.PostgresOutPort) *Factory {
	return &Factory{BaseUC{Out: out}}
}

func (f Factory) Query() QueryUsecase     { return QueryUsecase{BaseUC: f.BaseUC} }
func (f Factory) Command() CommandUsecase { return CommandUsecase{BaseUC: f.BaseUC} }

/// DI 관련 효율화 공부 좀 더해보기
