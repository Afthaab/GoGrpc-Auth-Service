package handler

import (
	"context"
	"net/http"

	"github.com/auth/service/pkg/domain"
	pb "github.com/auth/service/pkg/pb"
	useCase "github.com/auth/service/pkg/usecase/interface"
)

type UserHandler struct {
	UseCase    useCase.UserUseCase
	JwtUseCase useCase.JwtUseCase
	pb.AuthServiceServer
}

func NewUserHandler(useCase useCase.UserUseCase, jwtUseCase useCase.JwtUseCase) *UserHandler {
	return &UserHandler{
		UseCase:    useCase,
		JwtUseCase: jwtUseCase,
	}
}
func (h *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	_, err := h.UseCase.Register(user)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error",
		}, err
	}
	return &pb.RegisterResponse{
		Status: http.StatusOK,
		Error:  "nil",
	}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	userDetails, err := h.UseCase.Login(user)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error in Database",
		}, err
	}

	accessToken, err := h.JwtUseCase.GenerateAccessToken(int(userDetails.Id), userDetails.Email, "user")
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error in Generating JWT token",
		}, err
	}

	return &pb.LoginResponse{
		Status:      http.StatusOK,
		Accesstoken: accessToken,
	}, nil

}
