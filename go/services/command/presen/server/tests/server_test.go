package sever_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/grpc-tutorial/go/pkg/mysqlCon"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHelperPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "presen/serverパッケージのテスト")
}

//全テストスイートが実行される前に一度だけ実行される

var _ = BeforeSuite(func() {
	absPath, _ := filepath.Abs("../../../infra/sqlboiler/config/database.toml")
	os.Setenv("DATABASE_TOML_PATH", absPath)
	err := mysqlCon.DBConnect() // データベースに接続する
	Expect(err).NotTo(HaveOccurred(), "データベース接続が失敗したのでテストを中止します。")

})
