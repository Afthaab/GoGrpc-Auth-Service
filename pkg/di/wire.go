//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/auth/service/pkg/api"
	handler "github.com/auth/service/pkg/api/handler"
	config "github.com/auth/service/pkg/config"
	db "github.com/auth/service/pkg/db"
	repo "github.com/auth/service/pkg/repository"
	useCase "github.com/auth/service/pkg/usecase"
	"github.com/google/wire"
)

func InitApi(cfg config.Config) (*http.ServerHttp, error) {
	wire.Build(
		db.ConnectToDb,
		repo.NewUserRepo,
		useCase.NewUserUseCase,
		useCase.NewJWTUseCase,
		handler.NewUserHandler,
		http.NewServerHttp)

	return &http.ServerHttp{}, nil
}

//go run github.com/google/wire/cmd/wire
