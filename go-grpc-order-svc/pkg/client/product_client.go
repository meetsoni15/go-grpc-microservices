package client

import (
	"context"
	"fmt"
	"time"

	"github.com/meetsoni1511/go-grpc-order-svc/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient(url string) ProductServiceClient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	cc, err := grpc.DialContext(ctx, url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Could not connect GRPC server at %s Error %v:", url, err)
		panic(err)
	}

	return ProductServiceClient{
		Client: pb.NewProductServiceClient(cc),
	}
}

func (p *ProductServiceClient) GetOneProduct(productId int64) (*pb.GetProductResponse, error) {
	resp, err := p.Client.GetOneProduct(context.Background(), &pb.GetProductRequest{
		Id: productId,
	})

	return resp, err
}

func (p *ProductServiceClient) DecreaseStock(productId, orderId int64) (*pb.DecreaseStockResponse, error) {
	resp, err := p.Client.DecreaseStock(context.Background(), &pb.DecreaseStockRequest{
		Id:      productId,
		OrderId: orderId,
	})
	return resp, err
}
