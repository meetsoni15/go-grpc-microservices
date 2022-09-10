package main

import (
	"fmt"
	"log"
	"net"

	"github.com/meetsoni1511/go-grpc-auth-svc/pkg/config"
	"github.com/meetsoni1511/go-grpc-auth-svc/pkg/db"
	"github.com/meetsoni1511/go-grpc-auth-svc/pkg/pb"
	"github.com/meetsoni1511/go-grpc-auth-svc/pkg/services"
	"github.com/meetsoni1511/go-grpc-auth-svc/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	log.Println(conf)

	dbHandler := db.Init(conf.DBUrl)
	jwt := utils.JwtWrapper{
		SecretKey:       conf.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	listener, err := net.Listen("tcp", conf.Port)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	service := services.Server{
		DBHandler: *dbHandler,
		Jwt:       jwt,
	}
	fmt.Println("Auth Svc on", conf.Port)
	pb.RegisterAuthServiceServer(s, &service)
	if err := s.Serve(listener); err != nil {
		log.Fatalln("Failed to serve grpc:", err)
	}

}
