package handler

import (
	handlerInterface "github.com/abhinandpn/project-ecom/pkg/api/handler/interfaces"
	services "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
)

type OrderHandler struct {
	orderUseCase services.OrderUseCase
}

func NewOrderHandler(usecase services.OrderUseCase) handlerInterface.OrderHandler {
	return &OrderHandler{
		orderUseCase: usecase,
	}
}
