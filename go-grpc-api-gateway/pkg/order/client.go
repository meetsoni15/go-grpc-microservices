package order

import (
	"context"
	"fmt"
	"time"

	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/config"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	cc, err := grpc.DialContext(ctx, c.OrderSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Println("Could not connect:", err)
		panic(err)
	}

	return pb.NewOrderServiceClient(cc)
}
