package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
)

type Cartrepository interface {

	// cart
	FindCartByUserId(ctx context.Context, uid uint) (domain.Cart, error)

	// User
	CreateCart(ctx context.Context, uid uint) error          // create a empty cart for user
	Addtocart(ctx context.Context, pid uint, uid uint) error // add product to cart
	// UserCart(ctx context.Context, uid uint) (res.CartRes, error) // list the full cart total
	// RemovePrdFromCart

	// managment for cart
	UpdateCartHelp(ctx context.Context, uid uint, price float64) error
	FindCartIdByUserId(ctx context.Context, id uint) (uint, error)
	FindProductFromCart(ctx context.Context, cid, pids uint) (bool, error) // check wether the product exist in the cart
}
