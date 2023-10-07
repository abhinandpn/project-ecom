package interfaces

import (
	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type Cartrepository interface {
	// Find
	FindCartByUId(id uint) (domain.UserCart, error)
	FindCartInfoById(id uint) (domain.CartInfo, error)
	FindProductIntoCart(id, pfid uint) (bool, error)
	
	// CURD
	CreateUserCart(id uint) (domain.UserCart, error)
	CreateCartinfo(id uint) (domain.CartInfo, error)
	AddToCart(id, pfid, qty uint) error
	RemoveCart(id, pfid uint) error

	// View Cart
	ViewCart(id uint) ([]res.CartDisplay, error)
	CartInfo(id uint) (res.CartInfo, error)
	ViewCartProductInfoidByUid(id uint) ([]int, error)
	ViewCartQuantityidByUid(id uint) ([]int, error)
}
