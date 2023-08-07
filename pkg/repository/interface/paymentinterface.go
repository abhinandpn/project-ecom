package interfaces

import "github.com/abhinandpn/project-ecom/pkg/domain"

type PaymentRepository interface {
	ListPaymentMethods() ([]domain.PaymentMethod, error)
	FindPaymentMethodById(id uint) (domain.PaymentMethod, error)
	AddPaymentMethod(method string, status bool) error
	GetPaymentMethodByName(name string) (domain.PaymentMethod, error)
	DeletePaymentMethod(id uint) error
	GetPaymentMethodById(id uint) (domain.PaymentMethod, error)

	// -----------------------
	AddPaymentStatus(name string) error
	EditPaymentStatus(id uint, name string) (domain.PaymentStatus, error)
	DeltePaymentStatus(id uint) error
	GetPaymentStatusByName(name string) (domain.PaymentStatus, error)
	GetPaymentStatusById(id uint) (domain.PaymentStatus, error)
	GetAllPaymentStatus() ([]domain.PaymentStatus, error)
	// -----------------------
}
