package main

import (
	"fmt"
	"log"
	"net"

	"github.com/meetsoni1511/go-grpc-order-svc/pkg/client"
	"github.com/meetsoni1511/go-grpc-order-svc/pkg/config"
	"github.com/meetsoni1511/go-grpc-order-svc/pkg/db"
	"github.com/meetsoni1511/go-grpc-order-svc/pkg/pb"
	"github.com/meetsoni1511/go-grpc-order-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	dbHandler := db.Init(conf.DBUrl)
	productClientSvc := client.InitProductServiceClient(conf.ProductSvcUrl)
	service := services.Server{
		H:          dbHandler,
		ProductSvc: productClientSvc,
	}

	listener, err := net.Listen("tcp", conf.Port)
	if err != nil {
		log.Printf("Failed to listen at port %s with error %v\n", conf.Port, err)
		panic(err)
	}

	s := grpc.NewServer()
	fmt.Println("Order Svc on", conf.Port)
	pb.RegisterOrderServiceServer(s, &service)
	if err = s.Serve(listener); err != nil {
		log.Printf("Failed to serve GRPC at %s with error %v\n", conf.Port, err)
	}
}
