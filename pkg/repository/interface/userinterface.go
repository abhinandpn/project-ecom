package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type UserRepository interface {

	// Finding
	FindUser(ctx context.Context, user domain.Users) (domain.Users, error)                 // Find By Any filed
	FindUserByEmail(ctx context.Context, email string) (domain.Users, error)               // Find By Email
	FindUserByNumber(ctx context.Context, number string) (domain.Users, error)             // Find By Number
	FindUserById(ctx context.Context, id uint) (domain.Users, error)                       // Find By User Id
	FindUserByUserName(ctx context.Context, username string) (domain.Users, error)         // Find By UserName
	ListUsers(ctx context.Context, pagination req.PageNation) (res.ProductResponce, error) // List Full Users

	// opratios
	SaveUser(ctx context.Context, user domain.Users) (UserId uint, err error) // Create new User
	DeleteUser(ctx context.Context, id uint) error                            // Delete User
	UpdateUser(ctx context.Context, info domain.Users) (domain.Users, error)  // Update User Info

	// Address
	AddAddress(ctx context.Context, Uid uint, addres req.ReqAddress) error     // Add Address
	UpdateAddress(ctx context.Context, Uid uint, address req.ReqAddress) error // Update Address
	ListAllAddress(ctx context.Context, Uid uint) ([]res.ResAddress, error)    // Delete Address
	GetAddressByUid(uid uint) (domain.Address, error)                          // get address

	// wishlist
	CreateWishList(id uint) error
	AddToWishlistItem(wid, pfid uint) error
	RemoveFromWishListItem(wid, pfid uint) error
	FindWishListByUid(id uint) (domain.WishList, error)
	FindWishListItemByWid(id uint) (domain.WishListItems, error)
	FindProductFromWIshListItem(Wid, pfid uint) (bool, error)
	ViewWishList(uid uint, pagination req.PageNation) ([]res.ViewWishList, error)
}
