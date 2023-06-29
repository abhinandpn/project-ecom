package interfaces

import "github.com/abhinandpn/project-ecom/pkg/util/res"

type OrderUseCase interface {
	CreateUserOrder(id uint) error
	AddOrderInfo(uid, aid uint, cpid string, price float64, status string) (uint, error)
	AddOrderItems(oid, pfid, qty uint) error
	OrderByPfId(uid, id uint) error
	CartOrderAll(uid, pyid uint) error
	OrderStatus(id, oid uint) (res.OrderStatus, error)
}
