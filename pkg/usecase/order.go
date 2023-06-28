package usecase

import (
	"errors"

	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
)

type OrderUseCase struct {
	orderRepo   interfaces.OrderRepository
	userRepo    interfaces.UserRepository
	productRepo interfaces.ProductRepository
}

func NewOrderUseCase(OrderRepo interfaces.OrderRepository,
	UserRepo interfaces.UserRepository,
	ProductRepo interfaces.ProductRepository) service.OrderUseCase {
	return &OrderUseCase{orderRepo: OrderRepo,
		userRepo:    UserRepo,
		productRepo: ProductRepo}
}
func (or *OrderUseCase) CreateUserOrder(id uint) error {

	// check of exist
	order, err := or.orderRepo.FindUserOrderById(id)
	if err != nil {
		return err
	}
	// if its no create
	if order.Id != 0 {
		res := errors.New("user have user order")
		return res
	} else {
		err := or.orderRepo.CreateUserOrder(id)
		if err != nil {
			return err
		}
	}
	// response
	return nil
}

func (or *OrderUseCase) AddOrderInfo(uid, aid uint, cpid string, price float64, status string) (uint, error) {

	var OrderId uint
	// find order id
	UserOrder, err := or.orderRepo.FindUserOrderById(uid)
	if err != nil {
		return OrderId, err
	}
	// add info
	OrderId, err = or.orderRepo.AddOrderInfo(UserOrder.Id, aid, cpid, price, status)
	if err != nil {
		return OrderId, err
	}
	// response
	return OrderId, nil
}

func (or *OrderUseCase) AddOrderItems(oid, pfid, qty uint) error {

	err := or.orderRepo.AddOrderItem(oid, pfid, qty)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderUseCase) OrderByPfId(uid, id uint) error {

	// find the user order table
	userorder, err := o.orderRepo.FindUserOrderById(uid)
	if err != nil {
		return err
	}
	// find user address
	address, err := o.userRepo.GetAddressByUid(uid)
	if err != nil {
		return err
	}

	// find product info
	productinfo, err := o.productRepo.FindProductInfoById(id)
	if err != nil {
		return err
	}

	var cpid string
	var status string

	// add to orderinfo
	oid, err := o.orderRepo.AddOrderInfo(userorder.Id, address.ID, cpid, productinfo.Price, status)
	if err != nil {
		return err
	}

	// add to order items
	err = o.orderRepo.AddOrderItem(oid, productinfo.Id, 1)
	if err != nil {
		return err
	}

	// response
	return nil
}
