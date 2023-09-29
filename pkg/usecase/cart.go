package usecase

import (
	"errors"
	"fmt"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	services "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type CartUseCase struct {
	cartRepo interfaces.Cartrepository
	prd      interfaces.ProductRepository
	coupon   interfaces.Couponrepository
}

func NewCartUseCase(CartRepo interfaces.Cartrepository,
	p interfaces.ProductRepository,
	cp interfaces.Couponrepository) services.CartUseCase {

	return &CartUseCase{cartRepo: CartRepo,
		prd:    p,
		coupon: cp}
}
func (c *CartUseCase) FindCartInfoById(id uint) (domain.CartInfo, error) {

	var body domain.CartInfo
	// check if user have cart
	cart, err := c.cartRepo.FindCartByUId(id)
	if err != nil {
		return body, err
	}
	if cart.Id != 0 {
		body, err := c.cartRepo.FindCartInfoById(id)
		if err != nil {
			return body, err
		}
	} else {
		res := errors.New("user doest have cart info")
		return body, res
	}
	// if exist response
	return body, nil
}
func (c *CartUseCase) Createcart(id uint) error {

	// checking
	cart, err := c.cartRepo.FindCartByUId(id)
	if err != nil {
		return err
	}
	if cart.Id == 0 {
		cart, err = c.cartRepo.CreateUserCart(id)
		if err != nil {
			return err
		}
	} else {
		res := errors.New("user alredy have cart")
		return res
	}
	// if doest not exist create
	return nil
}

func (c *CartUseCase) AddToCart(id, pfid, qty uint) error {

	var cart domain.UserCart

	// check if user have cart
	cart, err := c.cartRepo.FindCartByUId(id)
	if err != nil {
		return err
	}

	// if does not exist create
	if cart.Id != 0 {
		exist, err := c.cartRepo.FindProductIntoCart(cart.Id, pfid)
		if err != nil {
			return err
		}
		if exist {
			res := errors.New("product alredy exist")
			return res
		} else {
			err := c.cartRepo.AddToCart(cart.Id, pfid, qty)
			if err != nil {
				return err
			}
		}
	}
	// response
	return nil
}

func (c *CartUseCase) RemoveFromCart(id, pfid uint) error {

	// check the usr have cart
	cart, err := c.cartRepo.FindCartByUId(id)
	if err != nil {
		return err
	}
	// check the product exist or not
	exist, err := c.cartRepo.FindProductIntoCart(cart.Id, pfid)
	if err != nil {
		return err
	}
	if !exist {
		res := errors.New("product does not exist")
		return res
	}
	if exist {
		err := c.cartRepo.RemoveCart(cart.Id, pfid)
		if err != nil {
			return err
		}
		if cart.CouponId != 0 {
			coupon, err := c.coupon.ViewCouponById(cart.CouponId)
			if err != nil {
				return err
			}
			newCartInfo, err := c.CartInfo(id)
			if err != nil {
				return err
			}
			if newCartInfo.Totalprice <= coupon.MinimumPurchase {
				err := c.coupon.RemoveCoupon(cart.UserId)
				if err != nil {
					return err
				}
			}
		}
	}
	// response
	return nil
}

func (c *CartUseCase) CartDisplay(id uint) ([]res.CartDisplay, error) {

	var body []res.CartDisplay
	// find if user have cart
	cart, err := c.cartRepo.FindCartByUId(id)
	if err != nil {
		return body, err
	}
	if cart.Id == 0 {
		res := errors.New("user cart doest not exst")
		return body, res
	}
	fmt.Println("cart in usecase  - >", cart)
	body, err = c.cartRepo.ViewCart(id)
	if err != nil {
		return body, err
	}
	return body, nil
}

func (c *CartUseCase) CartInfo(id uint) (res.CartInfo, error) {

	fmt.Println("uid --->", id)
	body, err := c.cartRepo.CartInfo(id)
	fmt.Println("body from usecase (cartindo -> boy)", body)
	if err != nil {
		return body, err
	}
	// --------
	cart, err := c.cartRepo.FindCartByUId(id)
	if err != nil {
		return body, err
	}

	coupon, err := c.coupon.ViewCouponById(cart.CouponId)
	if err != nil {
		return body, err
	}
	fmt.Println("------------>>>", coupon)
	// body updation
	body.CouponCode = coupon.Code
	body.DiscountPrice = coupon.DiscountPrice
	body.Totalprice = (body.Subtotal - body.DiscountPrice)
	// --------
	// if body.Subtotal == 0 {
	// 	res := errors.New("cart dosent have any product")
	// 	return body, res
	// }
	// response
	return body, nil
}

func (c *CartUseCase) CartInfoNew(id uint) (res.CartInfo, error) {

	var body res.CartInfo
	return body, nil
}
