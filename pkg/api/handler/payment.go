package handler

import (
	handlerInterface "github.com/abhinandpn/project-ecom/pkg/api/handler/interfaces"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
)

type PaymentHandler struct {
	PaymentUseCase service.PaymentuseCase
}

func NewPaymentHandler(paymentUseCase service.PaymentuseCase) handlerInterface.PaymentHandler {

	return &PaymentHandler{PaymentUseCase: paymentUseCase}
}
