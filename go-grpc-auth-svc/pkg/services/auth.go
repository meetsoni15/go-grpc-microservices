package services

import (
	"context"
	"net/http"

	"github.com/meetsoni1511/go-grpc-auth-svc/pkg/db"
	"github.com/meetsoni1511/go-grpc-auth-svc/pkg/models"
	"github.com/meetsoni1511/go-grpc-auth-svc/pkg/pb"
	"github.com/meetsoni1511/go-grpc-auth-svc/pkg/utils"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	DBHandler db.Handler
	Jwt       utils.JwtWrapper
}

func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	result := s.DBHandler.DB.Model(&user).Where(&models.User{Email: in.GetEmail()}).First(&user)
	if result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	if !utils.CheckPassword(in.GetPassword(), user.Password) {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  "Username or password is wrong",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if result := s.DBHandler.DB.Model(&user).Where(&models.User{Email: in.GetEmail()}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	user.Email = in.Email
	user.Password = utils.HashPassword(in.GetPassword())

	s.DBHandler.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Validate(ctx context.Context, in *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(in.GetToken())
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user = new(models.User)
	if result := s.DBHandler.DB.Model(&models.User{}).Where(&models.User{Email: claims.Email}).First(user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		User: &pb.UserData{
			UserId: user.ID,
			Email:  user.Email,
		},
	}, nil
}
