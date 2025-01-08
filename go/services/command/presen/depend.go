package presen

import (
	"github.com/grpc-tutorial/go/services/command/application"
	"github.com/grpc-tutorial/go/services/command/presen/prepare"
	"go.uber.org/fx"
)

var CommandDepend = fx.Options(
	application.SrvDepend,

	fx.Provide(),

	fx.Invoke(prepare.CommandServiceLifecycle),
)
