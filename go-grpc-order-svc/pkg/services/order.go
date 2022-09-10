package services

import (
	"context"
	"log"
	"net/http"

	"github.com/meetsoni1511/go-grpc-order-svc/pkg/client"
	"github.com/meetsoni1511/go-grpc-order-svc/pkg/db"
	"github.com/meetsoni1511/go-grpc-order-svc/pkg/models"
	"github.com/meetsoni1511/go-grpc-order-svc/pkg/pb"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedOrderServiceServer
	H          db.Handler
	ProductSvc client.ProductServiceClient
}

func (s *Server) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	productSvcResp, err := s.ProductSvc.GetOneProduct(in.Product)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("GRPC error %s\n", e.Message())
			log.Printf("GRPC error code %s\n", e.Code())
		} else {
			log.Printf("A non GRPC error %s\n", err)
		}
		return &pb.CreateOrderResponse{
			Status: productSvcResp.Status,
			Error:  productSvcResp.Error,
		}, nil
	}

	order := models.Order{
		Price:     productSvcResp.Data.Price,
		ProductId: productSvcResp.Data.Id,
		UserId:    in.UserId,
	}

	if result := s.H.DB.Create(&order); result.Error != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  result.Error.Error(),
		}, nil
	}

	_, err = s.ProductSvc.DecreaseStock(in.Product, order.Id)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("GRPC error %s\n", e.Message())
			log.Printf("GRPC error code %s\n", e.Code())
		} else {
			log.Printf("A non GRPC error %s\n", err)
		}
		return &pb.CreateOrderResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}
