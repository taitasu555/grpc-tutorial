package service

import (
	"context"

	categories "github.com/grpc-tutorial/go/services/command/domain/models/category"
)

// Category更新サービスインターフェイス
type CategoryService interface {
	// カテゴリを登録する
	Add(ctx context.Context, category *categories.Category) error
	// カテゴリを変更する
	Update(ctx context.Context, category *categories.Category) error
	// カテゴリを削除する
	Delete(ctx context.Context, categoryId *categories.Category) error
}
