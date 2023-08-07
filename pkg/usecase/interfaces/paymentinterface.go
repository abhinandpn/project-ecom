package interfaces

import "github.com/abhinandpn/project-ecom/pkg/domain"

type PaymentuseCase interface {
	PaymentMethods() ([]domain.PaymentMethod, error)
	AddPaymentMethod(name string, status bool) error
	DeletePaymentMethod(id uint) error

	// ------------------------
	CreatePaymentStatus(name string) error
	UpdatePaymentStatus(id uint, name string) (domain.PaymentStatus, error)
	DeltePaymentStatus(id uint) error
	FindPaymentStatusById(id uint) (domain.PaymentStatus, error)
	GetAllPaymentStatus() ([]domain.PaymentStatus, error)
	// ------------------------

}
