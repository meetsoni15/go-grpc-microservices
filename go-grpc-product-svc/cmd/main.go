package main

import (
	"fmt"
	"log"
	"net"

	"github.com/meetsoni1511/go-grpc-product-svc/pkg/config"
	"github.com/meetsoni1511/go-grpc-product-svc/pkg/db"
	"github.com/meetsoni1511/go-grpc-product-svc/pkg/pb"
	"github.com/meetsoni1511/go-grpc-product-svc/pkg/services"
	"google.golang.org/grpc"
)

func init() {

}

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	dbHandler := db.Init(conf.DBUrl)

	listener, err := net.Listen("tcp", conf.Port)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	fmt.Println("Product Svc on", conf.Port)
	service := services.Server{
		H: dbHandler,
	}
	pb.RegisterProductServiceServer(s, &service)
	if err := s.Serve(listener); err != nil {
		log.Fatalln("Failed to serve grpc:", err)
	}
}
