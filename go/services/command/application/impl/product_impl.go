package impl

import (
	"context"

	"github.com/grpc-tutorial/go/pkg/mysqlCon"
	"github.com/grpc-tutorial/go/services/command/application/service"
	"github.com/grpc-tutorial/go/services/command/domain/models/products"
)

type productServiceImpl struct {
	req products.ProductRepository
	mysqlCon.Transaction
}

func NewProductServiceImpl(
	req products.ProductRepository,
) service.ProductService {
	return &productServiceImpl{
		req: req,
	}
}

// 商品を登録する
func (ins *productServiceImpl) Add(ctx context.Context, product *products.Product) error {
	// トランザクションを開始する
	tran, err := ins.Begin(ctx)
	if err != nil {
		return err
	}
	// 実行結果に応じてトランザクションのコミットロールバック制御する
	defer func() {
		err = ins.Complete(tran, err)
	}()
	// 既に登録済であるか、確認する
	if err = ins.req.Exists(ctx, tran, product); err != nil {
		return err
	}
	// 商品を登録する
	if err = ins.req.Create(ctx, tran, product); err != nil {
		return err
	}
	return err
}

// カテゴリを変更する
func (ins *productServiceImpl) Update(ctx context.Context, product *products.Product) error {
	// トランザクションを開始する
	tran, err := ins.Begin(ctx)
	if err != nil {
		return err
	}
	// 実行結果に応じてトランザクションのコミットロールバック制御する
	defer func() {
		err = ins.Complete(tran, err)
	}()
	// 商品を変更する
	if err = ins.req.UpdateById(ctx, tran, product); err != nil {
		return err
	}
	return err
}

// 商品を削除する
func (ins *productServiceImpl) Delete(ctx context.Context, product *products.Product) error {
	// トランザクションを開始する
	tran, err := ins.Begin(ctx)
	if err != nil {
		return err
	}
	// 実行結果に応じてトランザクションのコミットロールバック制御する
	defer func() {
		err = ins.Complete(tran, err)
	}()
	// 商品を削除する
	if err = ins.req.DeleteById(ctx, tran, product); err != nil {
		return err
	}
	return err
}
