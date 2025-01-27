## リクエストの流れ
1. presen/server
2. adapter (パラメータと実行結果の変換インターフェス)
3. application/service サービス層のinterface
4. applicaiton/impl 処理の実態
5. domain/models(repository層のinterface)
6. repository(dbにアクセスしてデータを作成)

## テストフレームワーク
- Ginkgo (github.com/onsi/ginkgo/v2)
```
階層的なテスト構造: Describe, Context, It などのブロックを使用して、テストケースを階層的に構造化できます。
セットアップとティアダウン: BeforeEach, AfterEach, BeforeSuite, AfterSuite などのフックを使って、テスト前後の準備や後片付けを簡単に実装できます。
並列実行: テストの並列実行をサポートしており、大規模なテストスイートの実行時間を短縮できます。
カスタマイズ可能なレポート: 標準出力以外にも、様々な形式でテスト結果を出力できます。
```


-  Gomega (github.com/onsi/gomega)
```
Gomega は Ginkgo と組み合わせて使用される、強力なアサーションライブラリです。豊富なマッチャー（期待値）を提供しており、複雑な条件のチェックも簡潔に記述できます。
```
