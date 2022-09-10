package product

import (
	"context"
	"log"
	"time"

	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/config"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/product/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	cc, err := grpc.DialContext(ctx, c.ProductSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	log.Println("Starting Product GRPC server")

	return pb.NewProductServiceClient(cc)
}
