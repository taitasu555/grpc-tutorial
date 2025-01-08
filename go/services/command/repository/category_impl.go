package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/grpc-tutorial/go/pkg/mysqlCon"
	categories "github.com/grpc-tutorial/go/services/command/domain/models/category"
	"github.com/grpc-tutorial/go/services/command/infra/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type categoryRepositorySqlBoiler struct{}

func NewCategoryRepositorySqlBoiler() categories.CategoryRepository {

	//フック関数
	models.AddCategoryHook(boil.AfterInsertHook, categoryAfterInsertHook)
	models.AddCategoryHook(boil.AfterUpdateHook, categoryAfterUpdateHook)
	models.AddCategoryHook(boil.AfterDeleteHook, categoryAfterDeleteHook)
	return &categoryRepositorySqlBoiler{}
}

// 同名の商品カテゴリが存在確認結果を返す
func (ins *categoryRepositorySqlBoiler) Exists(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	// レコードの存在確認条件を作成する
	condition := models.CategoryWhere.Name.EQ(category.Name().Value())
	// レコードの存在を確認するクエリを実行する
	if exists, err := models.Categories(condition).Exists(ctx, tran); err != nil {
		return mysqlCon.DBErrHandler(err)
	} else if !exists { // 同じ名称のカテゴリは存在していない?
		return nil
	} else {
		return mysqlCon.DBErrHandler(fmt.Errorf("カテゴリ名:%sは既に登録されています。", category.Name().Value()))
	}
}

// 新しい商品カテゴリを永続化する
func (rep *categoryRepositorySqlBoiler) Create(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	// SqlBoilerのモデルを生成する
	new_category := models.Category{
		ID:    0,
		ObjID: category.Id().Value(),
		Name:  category.Name().Value(),
	}
	// 商品カテゴリを永続化する
	if err := new_category.Insert(ctx, tran, boil.Whitelist("obj_id", "name")); err != nil {
		return mysqlCon.DBErrHandler(err)
	}
	return nil
}

// 商品カテゴリを変更する
func (rep *categoryRepositorySqlBoiler) UpdateById(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	// 更新対象を取得する
	up_model, err := models.Categories(qm.Where("obj_id = ?", category.Id().Value())).One(ctx, tran)
	if up_model == nil {
		return mysqlCon.DBErrHandler(fmt.Errorf("カテゴリID:%sは存在しません。", category.Id().Value()))
	}
	if err != nil {
		return mysqlCon.DBErrHandler(err)
	}
	// 取得したモデルの値を変更する
	up_model.ObjID = category.Id().Value()
	up_model.Name = category.Name().Value()
	// 更新を実行する
	if _, err = up_model.Update(ctx, tran, boil.Whitelist("obj_id", "name")); err != nil {
		return mysqlCon.DBErrHandler(err)
	}
	return nil
}

// 商品カテゴリを削除する
func (rep *categoryRepositorySqlBoiler) DeleteById(ctx context.Context, tran *sql.Tx, category *categories.Category) error {
	// 削除対象を取得する
	del_model, err := models.Categories(qm.Where("obj_id = ?", category.Id().Value())).One(ctx, tran)
	if del_model == nil {
		return mysqlCon.DBErrHandler(fmt.Errorf("カテゴリID:%sは存在しません。", category.Id().Value()))
	}
	if err != nil {
		return mysqlCon.DBErrHandler(err)
	}
	// 削除を実行する
	if _, err = del_model.Delete(ctx, tran); err != nil {
		return mysqlCon.DBErrHandler(err)
	}
	return nil
}

// 登録処理後に実行されるフック
func categoryAfterInsertHook(ctx context.Context, exec boil.ContextExecutor, category *models.Category) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを登録しました。\n", category.ObjID, category.Name)
	return nil
}

// 変更処理後に実行されるフック
func categoryAfterUpdateHook(ctx context.Context, exec boil.ContextExecutor, category *models.Category) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを変更しました。\n", category.ObjID, category.Name)
	return nil
}

// 削除処理後に実行されるフック
func categoryAfterDeleteHook(ctx context.Context, exec boil.ContextExecutor, category *models.Category) error {
	log.Printf("カテゴリID:%s カテゴリ名:%sを削除しました。\n", category.ObjID, category.Name)
	return nil
}
