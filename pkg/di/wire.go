//go:build wireinject
// +build wireinject

package di

import (
	"github.com/amalmadhu06/simple-api/pkg/api/handler"
	"github.com/amalmadhu06/simple-api/pkg/config"
	"github.com/amalmadhu06/simple-api/pkg/db"
	"github.com/amalmadhu06/simple-api/pkg/repository"
	"github.com/amalmadhu06/simple-api/pkg/usecase"
	"github.com/google/wire"

	http "github.com/amalmadhu06/simple-api/pkg/api"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {

	wire.Build(
		db.ConnectDatabase,

		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,

		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
