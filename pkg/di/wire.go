//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/abhinandpn/project-ecom/pkg/api"
	handler "github.com/abhinandpn/project-ecom/pkg/api/handler"
	config "github.com/abhinandpn/project-ecom/pkg/config"
	db "github.com/abhinandpn/project-ecom/pkg/db"
	repository "github.com/abhinandpn/project-ecom/pkg/repository"
	usecase "github.com/abhinandpn/project-ecom/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
