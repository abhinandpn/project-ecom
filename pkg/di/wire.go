//go:build wireinject
// +build wireinject

package die

import (
	http "github.com/abhinandpn/project-ecom/pkg/api"
	"github.com/abhinandpn/project-ecom/pkg/api/handler"
	"github.com/abhinandpn/project-ecom/pkg/config"
	"github.com/abhinandpn/project-ecom/pkg/db"
	"github.com/abhinandpn/project-ecom/pkg/repository"
	"github.com/abhinandpn/project-ecom/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		// Database
		db.ConnectDatabase,
		// handler
		handler.NewAdminHandler,
		handler.NewUserHandler,
		handler.NewProductHandler,
		handler.NewCartHandler,
		handler.NewOrderHandler,
		handler.NewPaymentHandler,
		handler.NewCouponHandler,

		// usecase
		usecase.NewAdminUseCase,
		usecase.NewUserUseCase,
		usecase.NewCartUseCase,
		usecase.NewProductUseCase,
		usecase.NewOrderUseCase,
		usecase.NewPaymentUseCase,
		usecase.NewCouponUseCase,

		// repo
		repository.NewAdminRepository,
		repository.NewUserRepository,
		repository.NewProductRepository,
		repository.NewCartRepository,
		repository.NewOrderRepository,
		repository.NewPaymentRepository,
		repository.NewCouponRepository,

		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
