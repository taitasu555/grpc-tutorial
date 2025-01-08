package adapter

import (
	"github.com/grpc-tutorial/go/services/command/domain/models/products"
	"github.com/grpc-tutorial/go/services/command/proto/pb"
)

// パラメータと実行結果の変換インターフェス
type ProductAdapter interface {
	ToEntity(param *pb.ProductUpParam) (*products.Product, error)
	ToResult(result any) *pb.ProductUpResult
}
