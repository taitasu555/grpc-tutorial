package prepare

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/grpc-tutorial/go/pkg/mysqlCon"
	"go.uber.org/fx"
	"google.golang.org/grpc/reflection"
)

/*
/ 引数serverはfx.Provideで登録された値が渡る
*/
func CommandServiceLifecycle(lifecycle fx.Lifecycle, server *CommandServer) {
	lifecycle.Append(fx.Hook{
		// fx start
		OnStart: func(ctx context.Context) error {

			// データベース接続とコネクションプールを生成する
			if err := mysqlCon.DBConnect(); err != nil {
				panic(err)
			}

			port := 8082

			listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

			if err != nil {
				return err
			}

			reflection.Register(server.Server)

			//grpc serverを起動する
			go func() {
				log.Printf("start server port: %d", port)
				server.Server.Serve(listener)
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			// データベース接続を切断する
			return nil
		},
	},
	)
}
