package interfaces

import "github.com/abhinandpn/project-ecom/pkg/domain"

type PaymentRepository interface {
	FindPaymentMethodById(id uint) (domain.PaymentMethod, error)
	AddPaymentMethod(method string) error
}
