package interfaces

import (
	"context"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type UserUseCase interface {

	// Signup
	SignUp(ctx context.Context, user domain.Users) error
	Login(ctx context.Context, user domain.Users) (domain.Users, error)
	OtpLogin(ctx context.Context, user domain.Users) (domain.Users, error)

	UserAccount(ctx context.Context, UserId uint) (domain.Users, error)
	FindUserById(ctx context.Context, Uid uint) (domain.Users, error)

	// Address
	AddAddres(ctx context.Context, Uid uint, Address req.ReqAddress) error
	UpdateAddress(ctx context.Context, Uid uint, address req.ReqAddress) error
	ListAllAddress(ctx context.Context, Uid uint) ([]res.ResAddress, error)

	// wishlist
	FindWishList(id uint) (domain.WishList, error)
	FindWishLisItemByPFID(wid, pfid uint) (bool, error)
	CreteWishList(id uint) error
	AddToWishListItem(wid, pfid uint) error
	RemoveWishListItem(wid, pfid uint) error
	ViewWishList(uid uint, pagination req.PageNation) ([]res.ViewWishList, error)
}
