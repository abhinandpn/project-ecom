package handler

import (
	handlerInterface "github.com/abhinandpn/project-ecom/pkg/api/handler/interfaces"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
)

type CartsHandler struct {
	CartUseCase service.CartUseCase
}

func NewCartHandler(cartUseCase service.CartUseCase) handlerInterface.CartHandler {

	return &CartsHandler{CartUseCase: cartUseCase}
}
