package impl

import (
	"context"

	"github.com/grpc-tutorial/go/pkg/mysqlCon"

	"github.com/grpc-tutorial/go/services/command/application/service"
	categories "github.com/grpc-tutorial/go/services/command/domain/models/category"
)

type categoryServiceImpl struct {
	req categories.CategoryRepository
	mysqlCon.Transaction
}

func NewCategoryServiceImpl(
	req categories.CategoryRepository,
) service.CategoryService {
	return &categoryServiceImpl{req: req}
}

// カテゴリを登録する
func (ins *categoryServiceImpl) Add(ctx context.Context, category *categories.Category) error {
	tran, err := ins.Begin(ctx)
	if err != nil {
		return err
	}

	// 実行結果に応じてトランザクションのコミットロールバック制御する
	defer func() {
		defer ins.Complete(tran, err)
	}()

	if err = ins.req.Exists(ctx, tran, category); err != nil {
		return err
	}

	if err = ins.req.Create(ctx, tran, category); err != nil {
		return err
	}

	return nil
}

// カテゴリを変更する
func (ins *categoryServiceImpl) Update(ctx context.Context, category *categories.Category) error {
	// トランザクションの開始
	tran, err := ins.Begin(ctx)
	if err != nil {
		return err
	}
	// 実行結果に応じてトランザクションのコミットロールバック制御する
	defer func() {
		err = ins.Complete(tran, err)
	}()
	// カテゴリを更新する
	if err = ins.req.UpdateById(ctx, tran, category); err != nil {
		return err
	}
	return err
}

// カテゴリを削除する
func (ins *categoryServiceImpl) Delete(ctx context.Context, category *categories.Category) error {
	// トランザクションの開始
	tran, err := ins.Begin(ctx)
	if err != nil {
		return err
	}
	// 実行結果に応じてトランザクションのコミットロールバック制御する
	defer func() {
		err = ins.Complete(tran, err)
	}()
	// カテゴリを削除する
	if err = ins.req.DeleteById(ctx, tran, category); err != nil {
		return err
	}
	return err
}
