package usecase

import (
	"context"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	services "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
)

type CartUseCase struct {
	cartRepo interfaces.Cartrepository
	prd      interfaces.ProductRepository
}

func NewCartUseCase(CartRepo interfaces.Cartrepository, p interfaces.ProductRepository) services.CartUseCase {

	return &CartUseCase{cartRepo: CartRepo, prd: p}
}

// create a cart for new user with empty value

func (crt *CartUseCase) CreateCart(ctx context.Context, uid uint) (domain.Cart, error) {

	cart, err := crt.cartRepo.CreateCart(ctx, uid)
	if err != nil {
		return cart, err
	}
	return cart, nil
}

// find cart by user id

func (crt *CartUseCase) FindCartByUserId(ctx context.Context, uid uint) (domain.Cart, error) {

	body, err := crt.cartRepo.FindCartByUserId(ctx, uid)
	if err != nil {
		return body, err
	}
	return body, nil
}

func (crt *CartUseCase) Addtocart(ctx context.Context, cid, uid, pid, pfid uint) error {

	// check the usr have cart
	cart, err := crt.FindCartByUserId(ctx, uid)
	if err != nil {
		return err
	}

	// if doest have cart then create
	var body domain.Cart
	if cart.Id == 0 {
		body, err = crt.cartRepo.CreateCart(ctx, uid)
		if err != nil {
			return err
		}
	}

	// check the product exs in th store
	pinfo, err := crt.prd.FindProductInfoByPid(ctx, pid)
	if err != nil {
		return err
	}

	// check the qty
	if pinfo.Qty == 0 {
		return err
	}

	// add product in to cart
	cart, err = crt.cartRepo.Addtocart(ctx, body)
	if err != nil {
		return err
	}

	product, err := crt.prd.FindProductById(ctx, pid)
	if err != nil {
		return err
	}
	// update cart info
	err = crt.cartRepo.UpdateCartinfo(ctx, cid, 1, product.Price)
	if err != nil {
		return err
	}
	// return
	return err
}
