package usecase

import (
	"errors"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	services "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type CartUseCase struct {
	cartRepo interfaces.Cartrepository
	prd      interfaces.ProductRepository
}

func NewCartUseCase(CartRepo interfaces.Cartrepository, p interfaces.ProductRepository) services.CartUseCase {

	return &CartUseCase{cartRepo: CartRepo, prd: p}
}
func (c *CartUseCase) FindCartInfoById(id uint) (domain.CartInfo, error) {

	var body domain.CartInfo
	// check if user have cart
	cart, err := c.cartRepo.FindCartBy(id)
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
	cart, err := c.cartRepo.FindCartBy(id)
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

	
	var CartInfo domain.CartInfo
	// check if user have cart
	cart, err := c.cartRepo.FindCartBy(id)
	if err != nil {
		return err
	}

	// if does not exist create
	if cart.Id == 0 {
		cart, err = c.cartRepo.CreateUserCart(id)
		if err != nil {
			return err
		}
	}

	// check if user have cart info
	CartInfo, err = c.cartRepo.FindCartInfoById(cart.Id)
	if err != nil {
		return err
	}

	// if does not exist create one
	if CartInfo.Id == 0 {
		CartInfo, err = c.cartRepo.CreateCartinfo(cart.Id)
		if err != nil {
			return err
		}
	}

	// check if the product exist or not
	exist, err := c.cartRepo.FindProductIntoCart(CartInfo.Id, pfid)
	if err != nil {
		return err
	}

	if !exist {
		err := c.cartRepo.AddToCart(cart.Id, pfid, qty)
		if err != nil {
			return err
		}
	} else {
		res := errors.New("product alredy exist")
		return res
	}

	// response
	return nil
}

func (c *CartUseCase) RemoveFromCart(id, pfid uint) error {

	// check the usr have cart
	cart, err := c.cartRepo.FindCartBy(id)
	if err != nil {
		return err
	}
	// check the product exist or not
	exist, err := c.cartRepo.FindProductIntoCart(cart.Id, pfid)
	if err != nil {
		return err
	}
	if exist {
		err := c.cartRepo.RemoveCart(id, pfid)
		if err != nil {
			return err
		}
	}
	// response
	return nil
}

func (c *CartUseCase) CartDisplay(id uint) (res.CartDisplay, error) {

	var body res.CartDisplay
	// find if user have cart
	cart, err := c.cartRepo.FindCartBy(id)
	if err != nil {
		return body, err
	}
	if cart.Id == 0 {
		res := errors.New("user cart doest not exst")
		return body, res
	}
	body, err = c.cartRepo.ViewCart(id)
	if err != nil {
		return body, err
	}
	return body, nil
}
