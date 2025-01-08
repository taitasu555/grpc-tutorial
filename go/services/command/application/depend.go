package application

import (
	"github.com/grpc-tutorial/go/services/command/application/impl"
	"github.com/grpc-tutorial/go/services/command/repository"
	"go.uber.org/fx"
)

var SrvDepend = fx.Options(
	repository.ReqDepend,
	fx.Provide(
		impl.NewCategoryServiceImpl,
		impl.NewProductServiceImpl,
	),
)
