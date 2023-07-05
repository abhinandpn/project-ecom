package interfaces

import "github.com/abhinandpn/project-ecom/pkg/domain"

type PaymentuseCase interface {
	PaymentMethods() ([]domain.PaymentMethod, error)
	AddPaymentMethod(name string, status bool) error
	DeletePaymentMethod(id uint) error
}
