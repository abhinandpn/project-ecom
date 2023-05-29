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

func (crt *CartUseCase) AddToCart(ctx context.Context, pid, uid uint) error {

	err := crt.cartRepo.Addtocart(ctx, pid, uid)
	if err != nil {
		return err
	}

	product, err := crt.prd.FindProductById(ctx, pid)
	if err != nil {
		return err
	}

	productprice := product.Price

	err = crt.cartRepo.UpdateCartHelp(ctx, uid, productprice)

	if err != nil {
		return err
	}

	return nil
}
