package adapter

import (
	categories "github.com/grpc-tutorial/go/services/command/domain/models/category"
	"github.com/grpc-tutorial/go/services/command/proto/pb"
)

type CategoryAdapter interface {
	ToEntity(param *pb.CategoryUpParam) (*categories.Category, error)
	// 実行結果からCategoryUpResultに変換する
	ToResult(result any) *pb.CategoryUpResult
}
