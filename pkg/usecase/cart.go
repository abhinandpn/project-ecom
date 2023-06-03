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

func (crt *CartUseCase) CreateCart(ctx context.Context, uid uint) error {

	err := crt.cartRepo.CreateCart(ctx, uid)
	if err != nil {
		return err
	}
	return nil
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

	// check the user have cart and
	cart, err := crt.cartRepo.FindCartByUserId(ctx, uid)
	if err != nil {
		return err
	}

	// if user dosent have cart create
	if cart.Id == 0 {
		crt.cartRepo.CreateCart(ctx, uid)
	}

	// if have cart check user have cart info
	cartinfo, err := crt.cartRepo.FindCartInfoByCartId(ctx, cart.Id)
	if err != nil {
		return err
	}

	// if doest have cart info create
	if cartinfo.Id == 0 {
		crt.cartRepo.CreateCartInfo(ctx, cid)
	}

	// find product
	product, err := crt.prd.FindProductById(ctx, pid)
	if err != nil {
		return err
	}
	if product.Id == 0 {
		return err
	}
	// find prouct info
	pinfo, err := crt.prd.FindProductInfoByPid(ctx, pid)
	if err != nil {
		return err
	}
	// update quentity
	price := product.Price * float64(pinfo.Qty)

	// update cart info
	cartinfo, err = crt.cartRepo.UpdateCartinfo(ctx, cart.Id, pinfo.Qty, price)
	if err != nil {
		return err
	}
	// return
	return nil
}
