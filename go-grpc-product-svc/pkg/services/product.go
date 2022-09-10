package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/meetsoni1511/go-grpc-product-svc/pkg/db"
	"github.com/meetsoni1511/go-grpc-product-svc/pkg/models"
	"github.com/meetsoni1511/go-grpc-product-svc/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedProductServiceServer
	H db.Handler
}

func (s *Server) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product = models.Product{
		Name:  in.GetName(),
		Stock: in.GetStock(),
		Price: in.GetPrice(),
	}
	if result := s.H.DB.Model(&product).Create(&product); result.Error != nil {
		return &pb.CreateProductResponse{
				Status: http.StatusConflict,
				Error:  result.Error.Error(),
			}, status.Errorf(
				codes.AlreadyExists,
				result.Error.Error(),
			)
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.ID,
	}, nil
}

func (s *Server) GetOneProduct(ctx context.Context, in *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	var product = new(models.Product)
	if result := s.H.DB.First(product, in.Id); result.Error != nil {
		return &pb.GetProductResponse{
				Status: http.StatusNotFound,
				Error:  result.Error.Error(),
			}, status.Errorf(
				codes.NotFound,
				result.Error.Error(),
			)
	}
	return &pb.GetProductResponse{
		Status: http.StatusOK,
		Data: &pb.GetProductResponse_FindOneProduct{
			Id:    product.ID,
			Name:  product.Name,
			Stock: product.Stock,
			Price: product.Price,
		},
	}, nil
}

func (s *Server) DecreaseStock(ctx context.Context, in *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var (
		product = new(models.Product)
		log     = new(models.StockDecreaseLog)
	)

	tx := s.H.DB.Begin()

	if result := s.H.DB.First(product, in.Id); result.Error != nil {
		tx.Rollback()
		return &pb.DecreaseStockResponse{
				Status: http.StatusNotFound,
				Error:  result.Error.Error(),
			}, status.Errorf(
				codes.NotFound,
				result.Error.Error(),
			)
	}

	if product.Stock <= 0 {
		tx.Rollback()
		return &pb.DecreaseStockResponse{}, status.Errorf(
			codes.FailedPrecondition,
			fmt.Sprintln("product doesn't have specific stock"),
		)
	}

	//Check log exist for orderid
	result := s.H.DB.Model(log).Where(&models.StockDecreaseLog{Id: in.OrderId}).First(log)
	if result != nil {
		tx.Rollback()
		if result.Error != gorm.ErrRecordNotFound {
			return &pb.DecreaseStockResponse{
					Status: http.StatusInternalServerError,
					Error:  result.Error.Error(),
				}, status.Errorf(
					codes.Internal,
					result.Error.Error(),
				)
		}
	} else {
		var StockDecreasedErr = fmt.Sprintf("Stock already decreased for orderID %d", in.OrderId)
		return &pb.DecreaseStockResponse{
				Status: http.StatusConflict,
				Error:  StockDecreasedErr,
			}, status.Errorf(
				codes.AlreadyExists,
				StockDecreasedErr,
			)
	}

	product.Stock -= 1
	if result := s.H.DB.Save(product); result.Error != nil {
		tx.Rollback()
		return &pb.DecreaseStockResponse{
				Status: http.StatusInternalServerError,
				Error:  result.Error.Error(),
			}, status.Errorf(
				codes.Internal,
				result.Error.Error(),
			)
	}

	log.OrderId = in.GetOrderId()
	log.ProductRefer = product.ID

	if result := s.H.DB.Create(log); result.Error != nil {
		tx.Rollback()
		return &pb.DecreaseStockResponse{
				Status: http.StatusInternalServerError,
				Error:  result.Error.Error(),
			}, status.Errorf(
				codes.Internal,
				result.Error.Error(),
			)
	}

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}
