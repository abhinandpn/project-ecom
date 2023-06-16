package interfaces

import (
	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type Cartrepository interface {

	// Common Helper
	FindCartByUID(uid uint) (domain.Cart, error)
	FindCartInfoByCID(cid uint) (domain.CartInfo, error)
	FindProductFromCartByCId(pid uint) (domain.Cart, error)
	FindProductFromCartInfoByCId(pid uint) (domain.CartInfo, error)
	FindProductByPid(uid, pid uint) (bool, error)

	// create cart help
	CreateCartByUID(uid uint) (domain.Cart, error)
	CreateCartInfoByCid(cid uint) (domain.CartInfo, error)

	// add product to cart
	AddProductToCart(uid, pid, pfid uint) error
	AddProductToCartInfo(cid uint, pfr domain.Product) error

	// remove product
	RemoveProductfromCart(uid, pfid uint) error
	RemoveProductfromCartInfo(cid uint) error

	// list products
	ListAllProductFromCart(pagination req.PageNation, uid uint) ([]res.DisplayCart, error)
}
