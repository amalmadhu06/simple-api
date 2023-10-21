// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/amalmadhu06/simple-api/pkg/api"
	"github.com/amalmadhu06/simple-api/pkg/api/handler"
	"github.com/amalmadhu06/simple-api/pkg/config"
	"github.com/amalmadhu06/simple-api/pkg/db"
	"github.com/amalmadhu06/simple-api/pkg/repository"
	"github.com/amalmadhu06/simple-api/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	serverHTTP := http.NewServerHTTP(userHandler)
	return serverHTTP, nil
}
