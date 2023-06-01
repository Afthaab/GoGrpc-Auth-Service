// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/auth/service/pkg/api"
	"github.com/auth/service/pkg/api/handler"
	"github.com/auth/service/pkg/config"
	"github.com/auth/service/pkg/db"
	"github.com/auth/service/pkg/repository"
	"github.com/auth/service/pkg/usecase"
)

// Injectors from wire.go:

func InitApi(cfg config.Config) (*api.ServerHttp, error) {
	gormDB, err := db.ConnectToDb(cfg)
	if err != nil {
		return nil, err
	}
	userRepo := repository.NewUserRepo(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	jwtUseCase := usecase.NewJWTUseCase()
	userHandler := handler.NewUserHandler(userUseCase, jwtUseCase)
	serverHttp := api.NewServerHttp(userHandler)
	return serverHttp, nil
}