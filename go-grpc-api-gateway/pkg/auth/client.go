package auth

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*15)
	defer cancle()

	client, err := grpc.DialContext(ctx, c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	log.Println("Starting Auth GRPC server")

	return pb.NewAuthServiceClient(client)
}
