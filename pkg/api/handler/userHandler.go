package handler

import (
	"context"
	"errors"
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

// ------------------------------------- User Authentication ---------------------------------------------

func (h *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
	}
	err := h.UseCase.Register(user)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusUnauthorized,
			Error:  "Error in Registering the user",
		}, err
	}
	return &pb.RegisterResponse{
		Status: http.StatusOK,
		Error:  "nil",
	}, nil
}

func (h *UserHandler) RegisterValidate(ctx context.Context, req *pb.RegisterValidateRequest) (*pb.RegisterValidateResponse, error) {
	user := domain.User{
		Otp:   req.Otp,
		Email: req.Email,
	}
	user, err := h.UseCase.RegisterValidate(user)
	if err != nil {
		return &pb.RegisterValidateResponse{
			Status: http.StatusNotFound,
			Error:  "Error",
			Id:     int64(user.Id),
		}, err
	}
	return &pb.RegisterValidateResponse{
		Status: http.StatusOK,
		Error:  "nil",
		Id:     int64(user.Id),
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
			Status: http.StatusNotFound,
			Error:  "Error in logging the user",
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

func (u *UserHandler) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
	user := domain.User{
		Email: req.Email,
	}
	err := u.UseCase.ForgotPassword(user)
	if err != nil {
		return &pb.ForgotPasswordResponse{
			Status: http.StatusNotFound,
			Error:  "Error in Forget Passsword",
		}, err
	}
	return &pb.ForgotPasswordResponse{
		Status: http.StatusOK,
		Error:  "nil",
	}, nil

}
func (u *UserHandler) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ForgotPasswordResponse, error) {
	user := domain.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err := u.UseCase.ChangePassword(user)
	if err != nil {
		return &pb.ForgotPasswordResponse{
			Status: http.StatusNotFound,
			Error:  "Error in changing the password",
		}, err
	}
	return &pb.ForgotPasswordResponse{
		Status: http.StatusOK,
		Error:  "nil",
	}, nil
}

// -------------------------------------- Jwt Validation ---------------------------------------------

func (u *UserHandler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	userData := domain.User{}
	ok, claims := u.JwtUseCase.VerifyToken(req.Accesstoken)
	if !ok {
		return &pb.ValidateResponse{
			Status: http.StatusUnauthorized,
			Error:  "Token Verification Failed",
		}, errors.New("Token failed")
	}
	userData, err := u.UseCase.ValidateJwtUser(claims.Userid)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusUnauthorized,
			Userid: int64(userData.Id),
			Error:  "User not found with essesntial token credential",
			Source: claims.Source,
		}, err
	}
	return &pb.ValidateResponse{
		Status: http.StatusOK,
		Userid: int64(userData.Id),
		Source: claims.Source,
		Error:  "nil",
	}, nil

}

// ------------------------------------- Amdmin Authentication --------------------------------------------

func (h *UserHandler) AdminLogin(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user := domain.User{
		Username: req.Username,
		Password: req.Password,
	}
	userData, err := h.UseCase.AdminLogin(user)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "Error in Admin Login",
		}, err
	}
	accessToken, err := h.JwtUseCase.GenerateAccessToken(int(userData.Id), userData.Email, "admin")
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error in Generating JWT token",
		}, err
	}

	return &pb.LoginResponse{
		Status:      http.StatusOK,
		Accesstoken: accessToken,
		Error:       "nil",
	}, nil

}
