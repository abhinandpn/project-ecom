package interfaces

import (
	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type OrderRepository interface {
	CreateUserOrder(id uint) error
	FindUserOrderByUId(uid uint) (domain.UserOrder, error)
	AddOrderInfo(orderId, AdrsId, CopId uint, price float64, status uint, payid uint) (uint, error)
	FindOrderInfoByUid(uid uint) (domain.OrderInfo, error)
	AddOrderItem(oid, pfid, qty uint) error
	UpdatePaymentMethod(id, pid uint) error
	CartAllOrder(orderId, OrderinfoId uint, cart []res.CartDisplay) error // updation
	OrderDetail(uid uint) ([]res.ResOrder, error)
	ChangeOrderStatus(status string, id uint) error

	/*
		UPDATION NEED ....!
	*/
	FindOrderInfoByOrderId(id uint) (domain.OrderInfo, error)

	// Order Status
	AddOrderStatus(status string) error
	EditOrderStatus(id uint, status string) (domain.OrderStatus, error)
	DeleteOrderStatusById(id uint) error
	GetOrderStatusById(id uint) (domain.OrderStatus, error)
	GetOrderStatusByStatus(status string) (domain.OrderStatus, error)
	GetAllOrderStatus() ([]domain.OrderStatus, error)
	//
	// 01 - 09 - 2023 - Order status updation

	ListOrderDetailByUid(uid uint) ([]res.OrderDetailByUid, error)
	ListALlOrderByUid(id uint) ([]domain.OrderInfo, error)
	UpdateOrderStatusToOrdered(id uint) error
	UpdateOrderStatusToDelivered(id uint) error
	UpdateOrderStatusToCancelled(id uint) error
	UpdateOrderStatusToReturned(id uint) error
}
