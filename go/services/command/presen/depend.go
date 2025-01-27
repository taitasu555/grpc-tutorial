package presen

import (
	"github.com/grpc-tutorial/go/services/command/application"
	"github.com/grpc-tutorial/go/services/command/presen/adapter"
	"github.com/grpc-tutorial/go/services/command/presen/prepare"
	"github.com/grpc-tutorial/go/services/command/presen/server"
	"go.uber.org/fx"
)

/*
1. 依存関係の登録: fx.Provide によって各種アダプタやサーバーのコンストラクタが登録され、Fxがこれらを管理します。
2. 共通依存関係の追加: application.SrvDepend によって、アプリケーション全体で共有される依存関係が追加されます。
3. ライフサイクルの設定: fx.Invoke を通じて、アプリケーションのライフサイクル中に実行される関数が設定されます。
*/
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
