package usecase

import (
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
)

type PaymentUseCase struct {
	paymentRepo interfaces.PaymentRepository
}

func NewPaymentUseCase(PaymentRepo interfaces.PaymentRepository) service.PaymentuseCase {
	return &PaymentUseCase{paymentRepo: PaymentRepo}
}
