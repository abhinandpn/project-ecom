package interfaces

type OrderUseCase interface {
	CreateUserOrder(id uint) error
	AddOrderInfo(uid, aid uint, cpid string, price float64, status string) (uint, error)
	AddOrderItems(oid, pfid, qty uint) error
	OrderByPfId(uid, id uint) error
}
