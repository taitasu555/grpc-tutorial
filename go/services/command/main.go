package main

import (
	"github.com/grpc-tutorial/go/services/command/presen"
	"go.uber.org/fx"
)

func main() {
	// fxを起動する
	fx.New(
		presen.CommandDepend, // 依存性を定義する
	).Run()

}
