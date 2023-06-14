package usecase

import (
	"context"
	"errors"

	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	services "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type CartUseCase struct {
	cartRepo interfaces.Cartrepository
	prd      interfaces.ProductRepository
}

func NewCartUseCase(CartRepo interfaces.Cartrepository, p interfaces.ProductRepository) services.CartUseCase {

	return &CartUseCase{cartRepo: CartRepo, prd: p}
}

func (c *CartUseCase) AddProduct(uid, pfid uint) error {

	// check user have cart
	cart, err := c.cartRepo.FindCartByUID(uid)
	if err != nil {
		return err
	}

	// if doesnt have crete one
	if cart.Id == 0 {
		cart, err = c.cartRepo.CreateCartByUID(uid)
		if err != nil {
			return err
		}
	}

	product, err := c.prd.FindProductByPrinfo(pfid)
	if err != nil {
		return err
	}
	// check if porduct exist or not
	if cart.ProductId == product {
		return errors.New("product alredy exist")
	}

	// if dosent exist add
	err = c.cartRepo.AddProductToCart(uid, product, pfid)
	if err != nil {
		return err
	}

	// create cart info
	_, err = c.cartRepo.CreateCartInfoByCid(cart.Id)
	if err != nil {
		return err
	}

	var ctx context.Context
	prdt, err := c.prd.FindProductById(ctx, product)
	if err != nil {
		return err
	}
	// update cart info
	err = c.cartRepo.AddProductToCartInfo(cart.Id, prdt)
	if err != nil {
		return err
	}
	// response
	return nil
}

func (c *CartUseCase) CreateCart(uid uint) error {

	// check teh user have cart
	cart, err := c.cartRepo.FindCartByUID(uid)
	if err != nil {
		return err
	}

	// if its not exist create one
	if cart.Id == 0 {
		_, err := c.cartRepo.CreateCartByUID(uid)
		if err != nil {
			return err
		}
	}

	// return
	return nil
}

func (c *CartUseCase) RemoveProductFromCart(uid, pfid uint) error {

	// find the user cart
	cart, err := c.cartRepo.FindCartByUID(uid)
	if err != nil {
		return err
	}

	product, err := c.prd.FindProductByPrinfo(pfid)
	if err != nil {
		return err
	}
	// find the product exist or not
	ct, err := c.cartRepo.FindProductFromCartByCId(product)
	if err != nil {
		return err
	}

	if ct.Id != cart.Id {
		return err
	}

	// if exist remove
	err = c.cartRepo.RemoveProductfromCart(uid, pfid)
	if err != nil {
		return err
	}
	err = c.cartRepo.RemoveProductfromCartInfo(cart.Id)
	if err != nil {
		return err
	}
	// return
	return nil

}

func (c *CartUseCase) ListCart(id uint, pagination req.PageNation) ([]res.DisplayCart, error) {

	body, err := c.cartRepo.ListAllProductFromCart(pagination, id)
	if err != nil {
		return body, err
	}
	return body, nil
}
