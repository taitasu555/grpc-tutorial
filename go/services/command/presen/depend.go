package presen

import (
	"github.com/grpc-tutorial/go/services/command/application"
	"github.com/grpc-tutorial/go/services/command/presen/adapter"
	"github.com/grpc-tutorial/go/services/command/presen/prepare"
	"github.com/grpc-tutorial/go/services/command/presen/server"
	"go.uber.org/fx"
)

var CommandDepend = fx.Options(
	application.SrvDepend,
	fx.Provide(
		adapter.NewCategoryAdapterImpl,
		adapter.NewProductAdapterImpl,
		server.NewProductServer,
		server.NewCategoryServer,
		prepare.NewCommandServer,
	),

	fx.Invoke(prepare.CommandServiceLifecycle),
)
