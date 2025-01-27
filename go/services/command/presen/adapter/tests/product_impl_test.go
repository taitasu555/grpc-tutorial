package adapter_test

import (
	"github.com/grpc-tutorial/go/services/command/domain/errs"
	categories "github.com/grpc-tutorial/go/services/command/domain/models/category"
	"github.com/grpc-tutorial/go/services/command/domain/models/products"
	"github.com/grpc-tutorial/go/services/command/presen/adapter"
	"github.com/grpc-tutorial/go/services/command/proto/pb"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("productAdapterImpl構造体", Ordered, Label("メソッドのテスト"), func() {
	var category *categories.Category
	var adapt adapter.ProductAdapter

	BeforeAll(func() {
		adapt = adapter.NewProductAdapterImpl()

		id, _ := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
		category = categories.BuildCategory(id, nil)
	})

	Context("ToEntity()メソッド", func() {
		It(("Id,Name, Price, CateogyIdフィールドに値を設定する"), func() {
			param := pb.ProductUpParam{
				Crud:       pb.CRUD_UPDATE,
				Id:         "ac413f22-0cf1-490a-9635-7e9ca810e544",
				Name:       "水性ボールペン",
				Price:      120,
				CategoryId: "b1524011-b6af-417e-8bf2-f449dd58b5c0",
			}

			result, _ := adapt.ToEntity(&param)
			product_id, _ := products.NewProductId(param.Id)
			product_name, _ := products.NewProductName(param.GetName())
			product_price, _ := products.NewProductPrice(uint32(120))
			p := products.BuildProduct(product_id, product_name, product_price, category)
			Expect(result).To(Equal(p))
		})
		It("Name、Price、CategoryIdフィールドに値を設定する", func() {
			param := pb.ProductUpParam{
				Crud:       pb.CRUD_INSERT,
				Id:         "",
				Name:       "水性ボールペン(黒)",
				Price:      120,
				CategoryId: "b1524011-b6af-417e-8bf2-f449dd58b5c0",
			}
			result, _ := adapt.ToEntity(&param)
			Expect((result.Id())).NotTo(BeNil())
			Expect(result.Name().Value()).To(Equal("水性ボールペン(黒)"))
			Expect(result.Price().Value()).To(Equal(uint32(120)))
			Expect(result.Category().Id().Value()).To(Equal("b1524011-b6af-417e-8bf2-f449dd58b5c0"))
		})
	})

	Context("ToResult()メソッド", func() {
		It("ProductResultに変換する", func() {
			product_id, _ := products.NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e544")
			product_name, _ := products.NewProductName("水性ボールペン(黒)")
			product_price, _ := products.NewProductPrice(uint32(120))
			product := products.BuildProduct(product_id, product_name, product_price, category)
			result := adapt.ToResult(product)

			c := pb.Category{Id: "b1524011-b6af-417e-8bf2-f449dd58b5c0", Name: ""}
			p := pb.Product{Id: "ac413f22-0cf1-490a-9635-7e9ca810e544",
				Name: "水性ボールペン(黒)", Price: 120, Category: &c}
			Expect(result.Product).To(Equal(&p))
			Expect(result.Error).To(BeNil())
		})
		It("ErrorResultに変換する", func() {
			err := errs.NewDomainError("エラーが発生しました。")
			result := adapt.ToResult(err)
			Expect(result.Product).To(BeNil())
			Expect(result.Error.Message).To(Equal("エラーが発生しました。"))
		})
	})
})
