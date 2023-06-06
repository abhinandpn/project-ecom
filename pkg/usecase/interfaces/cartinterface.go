package interfaces

import (
	"context"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
)

type CartUseCase interface {

	// cart
	FindCartByUserId(ctx context.Context, uid uint) (domain.Cart, error)

	// create a cart for user with empty value
	CreateCart(ctx context.Context, uid uint) (domain.Cart, error)
	Addtocart(ctx context.Context, cid, uid, pid, pfid uint) error
}
