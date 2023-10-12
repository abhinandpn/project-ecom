package interfaces

import (
	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type CartUseCase interface {
	FindCartInfoById(id uint) (domain.CartInfo, error)
	Createcart(id uint) error
	AddToCart(id, pfid, qty uint) error
	RemoveFromCart(id, pfid uint) error
	CartDisplay(id uint) ([]res.CartDisplay, error)
	CartInfo(id uint) (res.CartInfo, error)
	CartInfoNew(id uint)(res.CartInfo,error)
}
																																			