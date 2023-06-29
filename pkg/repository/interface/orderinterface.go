package interfaces

import (
	"github.com/abhinandpn/project-ecom/pkg/domain"
)

type OrderRepository interface {
	CreateUserOrder(id uint) error
	FindUserOrderById(uid uint) (domain.UserOrder, error)
	AddOrderInfo(uid, aid uint, cpid string, price float64, status string) (uint, error)
	FindAllOrderInfoByUid(uid uint) (domain.OrderInfo, error)
	AddOrderItem(oid, pfid, qty uint) error
	AddOrderItemCartAll(oid, cid uint) error
	UpdatePaymentMethod(id, pid uint) error
	// OrderStatus(id, oid uint) (res.OrderStatus, error)
	// CartOrderAll(cid uint) error

	//
	/*
		UPDATION NEED ....!
	*/
	FindOrderInfoByOrderId(id uint) (domain.OrderInfo, error)
}
