package usecase

import (
	"context"
	"errors"

	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type OrderUseCase struct {
	orderRepo   interfaces.OrderRepository
	userRepo    interfaces.UserRepository
	productRepo interfaces.ProductRepository
	cartRepo    interfaces.Cartrepository
	paymentRepo interfaces.PaymentRepository
}

func NewOrderUseCase(OrderRepo interfaces.OrderRepository,
	UserRepo interfaces.UserRepository,
	ProductRepo interfaces.ProductRepository,
	CartRepo interfaces.Cartrepository,
	PaymentRepo interfaces.PaymentRepository,
) service.OrderUseCase {
	return &OrderUseCase{orderRepo: OrderRepo,
		userRepo:    UserRepo,
		productRepo: ProductRepo,
		cartRepo:    CartRepo,
		paymentRepo: PaymentRepo,
	}
}

var ctx context.Context

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

func (o *OrderUseCase) CartOrderAll(uid, pyid uint) error {

	// get user order id
	order, err := o.orderRepo.FindUserOrderById(uid)
	if err != nil {
		return err
	}

	// get cart id by uid
	cart, err := o.cartRepo.FindCartByUId(uid)
	if err != nil {
		return err
	}

	// get cartinfo
	cartinfo, err := o.cartRepo.CartInfo(uid)
	if err != nil {
		return err
	}

	// get address
	address, err := o.userRepo.GetAddressByUid(uid)
	if err != nil {
		return err
	}

	// select payment method
	payment, err := o.paymentRepo.FindPaymentMethodById(pyid)
	if err != nil {
		return err
	}

	var cpid string
	status := "product ordered"

	// update order infos
	oid, err := o.orderRepo.AddOrderInfo(order.Id, address.ID, cpid, cartinfo.Totalprice, status)
	if err != nil {
		return err
	}

	// update payment
	err = o.orderRepo.UpdatePaymentMethod(oid, payment.Id)
	if err != nil {
		return err
	}

	// update order items
	err = o.orderRepo.AddOrderItemCartAll(oid, cart.Id)
	if err != nil {
		return err
	}

	// response
	return nil
}

func (c *OrderUseCase) OrderStatus(id, oid uint) (res.OrderStatus, error) {

	var body res.OrderStatus

	user, err := c.userRepo.FindUserById(ctx, id)
	if err != nil {
		return body, err
	}
	body.Name = user.UserName // name update (1/6)

	addres, err := c.userRepo.GetAddressByUid(user.ID)
	if err != nil {
		return body, err
	}
	body.Adress = addres // adress update (2/6)

	orderinfo, err := c.orderRepo.FindOrderInfoByOrderId(oid)
	if err != nil {
		return body, err
	}
	body.OrderStatus = orderinfo.OrderStatus // order status update (3/6)

	cart, err := c.cartRepo.ViewCart(user.ID)
	if err != nil {
		return body, err
	}
	body.CartRes = cart // updated

	return body, nil
}
