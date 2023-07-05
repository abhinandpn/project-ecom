package interfaces

import "github.com/abhinandpn/project-ecom/pkg/domain"

type PaymentRepository interface {
	ListPaymentMethods() ([]domain.PaymentMethod, error)
	FindPaymentMethodById(id uint) (domain.PaymentMethod, error)
	AddPaymentMethod(method string, status bool) error
	GetPaymentMethodByName(name string) (domain.PaymentMethod, error)
	DeletePaymentMethod(id uint) error
	GetPaymentMethodById(id uint) (domain.PaymentMethod, error)
}
