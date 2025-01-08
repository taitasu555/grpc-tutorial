package mysqlCon_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	mysqlCon "github.com/grpc-tutorial/go/pkg/mysqlCon"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "infra/sqlboiler/handlerパッケージのテスト")
}

var _ = Describe("データベース接続テスト", func() {
	It("接続が成功した場合、nilが返る", Label("DB接続"), func() {
		// 現在のファイルのディレクトリを取得
		_, filename, _, ok := runtime.Caller(0)
		Expect(ok).To(BeTrue(), "runtime.Caller failed")

		// テストファイルのディレクトリを基にdatabase.tomlのパスを解決
		dir := filepath.Dir(filename)
		absPath := filepath.Join(dir, "database.toml")

		// 環境変数に設定
		os.Setenv("DATABSE_TOML_PATH", absPath)
		defer os.Unsetenv("DATABSE_TOML_PATH") // テスト後に環境変数をクリーンアップ

		result := mysqlCon.DBConnect()
		Expect(result).To(BeNil())
	})
})
