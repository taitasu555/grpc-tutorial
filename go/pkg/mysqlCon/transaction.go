package mysqlCon

import (
	"context"
	"database/sql"
	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Transaction struct{}

// トランザクションを開始する
func (inc *Transaction) Begin(ctx context.Context) (*sql.Tx, error) {
	tran, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, DBErrHandler(err)
	}
	return tran, nil
}

func (inc *Transaction) Complete(tran *sql.Tx, err error) error {
	if err != nil {
		if e := tran.Rollback(); e != nil {
			return DBErrHandler(e)
		} else {
			log.Println("トランザクションをロールバックしました。")
		}
	} else {
		if e := tran.Commit(); e != nil {
			return DBErrHandler(e)
		} else {
			log.Println("トランザクションをコミットしました。")
		}
	}
	return nil
}
