package usecase

import (
	"context"
	"errors"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
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
	order, err := or.orderRepo.FindUserOrderByUId(id)
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

func (or *OrderUseCase) AddOrderItems(oid, pfid, qty uint) error {

	err := or.orderRepo.AddOrderItem(oid, pfid, qty)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderUseCase) OrderByPfId(uid, id uint) error {

	//
	//
	// UPDATION NEED
	//
	//
	//
	// response
	return nil
}

func (o *OrderUseCase) CartOrderAll(uid, payid, copId, adrsId uint) error {

	// get user orderid
	order, err := o.orderRepo.FindUserOrderByUId(uid)
	if err != nil {
		return err
	}
	// get user cart info
	// cart, err := o.cartRepo.FindCartByUId(uid)
	// if err != nil {
	// 	return err
	// }

	// get cartitems
	cartitem, err := o.cartRepo.ViewCart(uid)
	if err != nil {
		return err
	}
	if cartitem == nil {
		res := errors.New("cart is empty")
		return res
	}

	cartinfo, err := o.cartRepo.CartInfo(uid)
	if err != nil {
		return err
	}

	// get address
	address, err := o.userRepo.GetAddressByAdrsId(uid, adrsId)
	if err != nil {
		return err
	}
	// select payment method
	payment, err := o.paymentRepo.FindPaymentMethodById(payid)
	if err != nil {
		return err
	}
	status := "ordered"
	var orderInfo uint
	if address.ID == 0 {
		// get default address
		defaultad, err := o.userRepo.GetUserDefaultAddressId(uid)
		if err != nil {
			return err
		}
		if defaultad.ID == 0 {
			res := errors.New("no default address is found")
			return res
		}
		orderInfo, err = o.orderRepo.AddOrderInfo(order.Id,
			defaultad.ID, copId, cartinfo.Totalprice, status, payment.Id)
		if err != nil {
			return err
		}
	} else {
		// update order infos
		orderInfo, err = o.orderRepo.AddOrderInfo(order.Id,
			address.ID, copId, cartinfo.Totalprice, status, payment.Id)
		if err != nil {
			return err
		}
	}
	// if address.ID == 0 {
	// 	defaultad, err := o.userRepo.GetUserDefaultAddressId(uid)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	orderInfo, err = o.orderRepo.AddOrderInfo(order.Id,
	// 		defaultad.ID, copId, cartinfo.Totalprice, status, payment.Id)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	err = o.orderRepo.CartAllOrder(order.Id, orderInfo, cartitem)
	if err != nil {
		return err
	}
	//

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

func (o *OrderUseCase) UserOrders(id uint) ([]res.ResOrder, error) {

	body, err := o.orderRepo.OrderDetail(id)
	if err != nil {
		return body, err
	}
	if body == nil {
		res := errors.New("user does not have orders")
		return body, res
	}
	return body, nil
}

func (o *OrderUseCase) ChangeOrderStatus(status string, id uint) error {

	order, err := o.orderRepo.FindUserOrderByUId(id)
	if err != nil {
		return err
	}
	err = o.orderRepo.ChangeOrderStatus(status, order.Id)
	if err != nil {
		return err
	}
	return nil
}

// ------------- order status -------------

func (or *OrderUseCase) CreateOrderStatus(status string) error {

	body, err := or.orderRepo.GetOrderStatusByStatus(status)
	if err != nil {
		return err
	}
	if body.Id != 0 {
		res := errors.New("status alredy exist")
		return res
	}
	err = or.orderRepo.AddOrderStatus(status)
	if err != nil {
		return err
	}
	return nil
}

func (or *OrderUseCase) UpdateOrderStatus(id uint, status string) (domain.OrderStatus, error) {

	body, err := or.orderRepo.GetOrderStatusById(id)
	if err != nil {
		return body, err
	}
	if body.Id != 0 {
		order, err := or.orderRepo.GetOrderStatusByStatus(status)
		if err != nil {
			return order, err
		}
		if order.Id != 0 {
			res := errors.New("status alredy exist")
			return order, res
		}
		body, err = or.orderRepo.EditOrderStatus(id, status)
		if err != nil {
			return body, err
		}
	}
	return body, nil
}

func (or *OrderUseCase) DeletOrderStatus(id uint) error {

	body, err := or.orderRepo.GetOrderStatusById(id)
	if err != nil {
		return err
	}
	if body.Id != 0 {
		err := or.orderRepo.DeleteOrderStatusById(id)
		if err != nil {
			return err
		}
	} else {
		res := errors.New("order status does not exist")
		return res
	}
	return nil
}

func (or *OrderUseCase) FindOrderStatusById(id uint) (domain.OrderStatus, error) {

	body, err := or.orderRepo.GetOrderStatusById(id)
	if err != nil {
		return body, err
	}
	if body.Id == 0 {
		res := errors.New("order status does not exist")
		return body, res
	}
	return body, nil
}

func (or *OrderUseCase) FindOrderStatusByStatus(status string) (domain.OrderStatus, error) {

	body, err := or.orderRepo.GetOrderStatusByStatus(status)
	if err != nil {
		return body, err
	}
	if body.Id == 0 {
		res := errors.New("order status does not exist")
		return body, res
	}
	return body, nil
}

func (or *OrderUseCase) FindAllOrderStatus() ([]domain.OrderStatus, error) {

	body, err := or.orderRepo.GetAllOrderStatus()
	if err != nil {
		return body, err
	}
	if body == nil {
		res := errors.New("order status does not exist")
		return body, res
	}
	return body, nil
}
