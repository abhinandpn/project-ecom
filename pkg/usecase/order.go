package usecase

import (
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
)

type OrderUseCase struct {
	orderRepo interfaces.OrderRepository
}

func NewOrderUseCase(OrderRepo interfaces.OrderRepository) service.OrderUseCase {
	return &OrderUseCase{orderRepo: OrderRepo}
}
