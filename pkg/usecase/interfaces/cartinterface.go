package interfaces

import (
	"context"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type CartUseCase interface {

	// cart
	FindCartByUserId(ctx context.Context, uid uint) (domain.Cart, error)

	// create a cart for user with empty value
	CreateCart(ctx context.Context, uid uint) error
	AddToCart(ctx context.Context, pid, uid uint) error
	UserCart(ctx context.Context, uid uint) (res.CartRes, error)
}
