package mysqlCon

import (
	"errors"
	"log"
	"net"

	"github.com/go-sql-driver/mysql"
)

// データベースアクセスエラーのハンドリング
func DBErrHandler(err error) error {
	var opErr *net.OpError
	var driverErr *mysql.MySQLError
	if errors.As(err, &opErr) { // 接続がタイムアウトかネットワーク関連の問題が原因で接続が確立できない?
		log.Println(err.Error())
		return NewInternalError(opErr.Error())
	} else if errors.As(err, &driverErr) { // MySQLドライバエラー?
		log.Printf("Code:%d Message:%s \n", driverErr.Number, driverErr.Message)
		if driverErr.Number == 1062 { // 一意制約違反?
			return NewInternalError("一意制約違反です。")
		} else {
			return NewInternalError(driverErr.Message)
		}
	} else { // その他のエラー
		log.Println(err.Error())
		return NewInternalError(err.Error())
	}
}

type InternalError struct {
	message string // エラーメッセージ
}

// エラーメッセージを返すメソッド
func (e *InternalError) Error() string {
	return e.message
}

// コンストラクタ
func NewInternalError(message string) *InternalError {
	return &InternalError{message: message}
}
