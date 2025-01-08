package server

import (
	"context"

	"github.com/grpc-tutorial/go/services/command/application/service"
	"github.com/grpc-tutorial/go/services/command/presen/adapter"
	"github.com/grpc-tutorial/go/services/command/proto/pb"
)

type productServer struct {
	pb.UnimplementedProductCommandServer
	adapter adapter.ProductAdapter
	service service.ProductService
}

func NewProductServer(adapter adapter.ProductAdapter, service service.ProductService) pb.ProductCommandServer {
	return &productServer{
		adapter: adapter,
		service: service,
	}
}

func (ins *productServer) Create(ctx context.Context, param *pb.ProductUpParam) (*pb.ProductUpResult, error) {
	product, err := ins.adapter.ToEntity(param)
	if err != nil {
		return ins.adapter.ToResult(nil), err
	}

	if err := ins.service.Add(ctx, product); err != nil {
		return ins.adapter.ToResult(nil), err
	}

	return ins.adapter.ToResult(nil), nil
}

func (ins *productServer) Update(ctx context.Context, param *pb.ProductUpParam) (*pb.ProductUpResult, error) {
	product, err := ins.adapter.ToEntity(param)
	if err != nil {
		return ins.adapter.ToResult(err), nil
	}
	if err := ins.service.Update(ctx, product); err != nil {
		return ins.adapter.ToResult(err), nil
	}
	return ins.adapter.ToResult(product), nil
}

func (ins *productServer) Delete(ctx context.Context, param *pb.ProductUpParam) (*pb.ProductUpResult, error) {
	product, err := ins.adapter.ToEntity(param)
	if err != nil {
		return ins.adapter.ToResult(err), nil
	}
	if err := ins.service.Delete(ctx, product); err != nil {
		return ins.adapter.ToResult(err), nil
	}
	return ins.adapter.ToResult(product), nil
}
