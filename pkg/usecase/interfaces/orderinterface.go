package interfaces

import (
	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type OrderUseCase interface {
	CreateUserOrder(id uint) error
	AddOrderItems(oid, pfid, qty uint) error
	OrderByPfId(uid, id uint) error
	OrderStatus(id, oid uint) (res.OrderStatus, error)
	CartOrderAll(uid, payid, copId, adrsId uint) error // updation
	UserOrders(id uint) ([]res.ResOrder, error)
	ChangeOrderStatus(status string, id uint) error

	// order status
	CreateOrderStatus(status string) error
	UpdateOrderStatus(id uint, status string) (domain.OrderStatus, error)
	DeletOrderStatus(id uint) error
	FindOrderStatusById(id uint) (domain.OrderStatus, error)
	FindOrderStatusByStatus(status string) (domain.OrderStatus, error)
	FindAllOrderStatus() ([]domain.OrderStatus, error)

	UpdatedCartAllOrder(uid, payid, addid uint) error
	//
}
