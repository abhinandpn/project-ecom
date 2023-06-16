package interfaces

import (
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type CartUseCase interface {

	// Create cart
	// CreateCart(uid uint) error

	// Add Product from cart
	// AddProduct(uid, pfid uint) error

	// Remove product from cart
	RemoveProductFromCart(uid, pid uint) error

	// List product from cart
	ListCart(id uint, pagination req.PageNation) ([]res.DisplayCart, error)
}
