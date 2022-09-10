package main

import (
	"go-grpc-project/go-grpc-api-gateway/pkg/order"
	"log"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-swagger/go-swagger"
	_ "github.com/meetsoni1511/go-grpc-api-gateway/docs"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/auth"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/config"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/product"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	authSvc := auth.RegisterRoutes(r, &conf)
	product.RegisterRoutes(r, &conf, authSvc)
	order.RegisterRoutes(r, &conf, authSvc)

	log.Println("gateway listening at ", conf.Port)
	r.Run(conf.Port)
}
