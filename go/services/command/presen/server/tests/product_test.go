package sever_test

import (
	"context"

	"github.com/grpc-tutorial/go/services/command/application"
	"github.com/grpc-tutorial/go/services/command/presen/adapter"
	"github.com/grpc-tutorial/go/services/command/presen/server"
	"github.com/grpc-tutorial/go/services/command/proto/pb"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
)

var _ = Describe(("Product構造体のテスト"), Ordered, Label("メソッドのテスト"), func() {
	var srv pb.ProductCommandServer
	var product *pb.Product

	var ctx context.Context
	var container *fx.App

	BeforeAll(func() {
		ctx = context.Background()
		container = fx.New(
			application.SrvDepend,
			fx.Provide(
				adapter.NewProductAdapterImpl,
				server.NewProductServer,
			),
			fx.Populate(&srv),
		)

		err := container.Start(ctx)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		err := container.Stop(context.Background())
		Expect(err).NotTo(HaveOccurred())
	})

	Context("Createメソッドのテスト", Label("Createメソッド"), func() {
		It("正常系", func() {
			param := pb.ProductUpParam{
				Id:         "ac413f22-0cf1-490a-9635-7e9ca810e544",
				Crud:       pb.CRUD_INSERT,
				Name:       "水性ボールペン",
				Price:      100,
				CategoryId: "b1524011-b6af-417e-8bf2-f449dd58b5c0",
			}

			result, err := srv.Create(ctx, &param)

			product = result.Product
			Expect(err).NotTo(HaveOccurred())
			Expect(product).NotTo(BeNil())
		})
	})
})
