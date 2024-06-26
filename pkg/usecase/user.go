package usecase

import (
	"context"
	"errors"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/helper"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo  interfaces.UserRepository
	cartRepo  interfaces.Cartrepository
	orderRepo interfaces.OrderRepository
}

func NewUserUseCase(repo interfaces.UserRepository,
	CartRepo interfaces.Cartrepository,
	OrderRepo interfaces.OrderRepository) service.UserUseCase {
	return &userUseCase{userRepo: repo,
		cartRepo:  CartRepo,
		orderRepo: OrderRepo}
}

// ........................................
func (usr *userUseCase) SignUp(ctx context.Context, user domain.Users) error {
	// check alredy exist or not
	checkUser, err := usr.userRepo.FindUser(ctx, user)

	if err != nil {
		return err
	}
	// if user not exist create user
	if checkUser.ID == 0 {

		//------- hash password
		hashpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

		if err != nil {
			return err
		}

		// save the password in hash
		user.Password = string(hashpass)

		// save the user
		Uid, err := usr.userRepo.SaveUser(ctx, user)
		if err != nil {
			return err
		}
		err = usr.userRepo.CreateWishList(Uid) // wishlist creating
		if err != nil {
			return err
		}
		_, err = usr.cartRepo.CreateUserCart(Uid) // cart creating
		if err != nil {
			return err
		}
		err = usr.orderRepo.CreateUserOrder(Uid) // order table
		if err != nil {
			return err
		}
		return nil
	}

	// if user exist then check which field is exist
	return helper.UserCheck(user, checkUser)

}

// Login wiht OTP
func (usr *userUseCase) OtpLogin(ctx context.Context, user domain.Users) (domain.Users, error) {

	user, err := usr.userRepo.FindUser(ctx, user) // Find User from database

	if err != nil {
		return user, errors.New("can't find the user")
	} else if user.ID == 0 {
		return user, errors.New("user not exist with this details")
	}

	// Chech the user Block Status

	if user.IsBlocked {
		return user, errors.New("user blocked by admin")
	}

	return user, nil

}

//

func (usr *userUseCase) Login(ctx context.Context, user domain.Users) (domain.Users, error) {

	dbUser, dbErr := usr.userRepo.FindUser(ctx, user)

	if dbErr != nil {
		return user, dbErr
	} else if dbUser.ID == 0 {
		return user, errors.New("user not exist with this details")
	}

	// check the block status

	if dbUser.IsBlocked {
		return user, errors.New("user blocked by admin")
	}

	// check username with password

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))

	if err != nil {
		// return user, errors.New("entered password is wrong")
		return user, err
	}

	return dbUser, nil
}

// User account Info
func (usr *userUseCase) UserAccount(ctx context.Context, UserId uint) (domain.Users, error) {

	var user domain.Users
	user, err := usr.userRepo.FindUserById(ctx, UserId)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (usr *userUseCase) FindUserById(ctx context.Context, Uid uint) (domain.Users, error) {
	var user domain.Users
	user, err := usr.userRepo.FindUserById(ctx, Uid)
	return user, err
}

// -----------------AddAddress-----------------

func (usr *userUseCase) AddAddres(ctx context.Context, Uid uint, Address req.ReqAddress) error {

	body1, err := usr.userRepo.GetAddressByName(Address.Name, Uid)
	if err != nil {
		return err
	}
	if body1.ID != 0 {
		body2, err := usr.userRepo.GetAddressByHouseName(Address.House, Uid)
		if err != nil {
			return err
		}
		if body2.ID != 0 {
			body3, err := usr.userRepo.GetAddressByNumber(Address.PhoneNumber, Uid)
			if err != nil {
				return err
			}
			if body3.ID != 0 {
				body4, err := usr.userRepo.GetAddressByPinCode(Address.Pincode, Uid)
				if err != nil {
					return err
				}
				if body4.ID != 0 {
					res := errors.New("address alredy exist with this details")
					return res
				}
			}
		}
	}

	err = usr.userRepo.AddAddress(ctx, Uid, Address)
	if err != nil {
		return err
	}

	return err

}

func (usr *userUseCase) UpdateAddress(ctx context.Context, Uid uint, address req.ReqAddress) error {

	var body req.ReqAddress

	err := usr.userRepo.UpdateAddress(ctx, Uid, body)

	return err
}
func (usr *userUseCase) ListAllAddress(Uid uint) ([]res.ResAddress, error) {

	var body []res.ResAddress

	body, err := usr.userRepo.ListAllAddress(Uid)
	if body == nil {
		res := errors.New("address is empty")
		return body, res
	}

	return body, err
}

func (usr *userUseCase) MakeAddressDefault(uid, id uint) error {

	// if user have address
	address, err := usr.userRepo.ListAllAddress(uid)
	if err != nil {
		return err
	}
	if address == nil {
		res := errors.New("user does not have address")
		return res
	}
	// if user have default address
	dflt, err := usr.userRepo.FindDefaultAddress(uid)
	if err != nil {
		return err
	}
	if dflt.ID != 0 {

		// if have change the default address to undefault
		err := usr.userRepo.AddressRemoveDefaultById(dflt.ID)
		if err != nil {
			return err
		}
		// make new address id default
		err = usr.userRepo.MakeAddressDefaultById(id)
		if err != nil {
			return err
		}
	}
	err = usr.userRepo.MakeAddressDefaultById(id)
	if err != nil {
		return err
	}

	return nil
}

func (a *userUseCase) GetUserDefaultAddressId(uid uint) (domain.Address, error) {

	var body domain.Address

	address, err := a.userRepo.ListAllAddress(uid)
	if err != nil {
		return body, err
	}
	if address == nil {
		res := errors.New("address is empty")
		return body, res
	}
	body, err = a.userRepo.GetUserDefaultAddressId(uid)
	if err != nil {
		return body, err
	}
	return body, nil
}

// ----------------------- wishlist -------------------------------

func (w *userUseCase) FindWishList(id uint) (domain.WishList, error) {

	var body domain.WishList
	body, err := w.userRepo.FindWishListByUid(id)

	if err != nil {
		return body, err
	}
	return body, nil
}

func (w *userUseCase) FindWishLisItemByPFID(wid, pfid uint) (bool, error) {

	var body bool

	wishlist, err := w.userRepo.FindWishListItemByWid(wid)
	if err != nil {
		if err != nil {
			return body, err
		}
	}
	// fmt.Println("wishlist id ", wishlist.Id)
	// if wishlist.Id == 0 {
	// 	res := errors.New("user does have wishlist item")
	// 	return body, res
	// }

	body, err = w.userRepo.FindProductFromWIshListItem(wishlist.WishListId, pfid)

	if err != nil {
		return body, err
	}
	return body, nil
}

func (w *userUseCase) CreteWishList(id uint) error {

	wishlist, err := w.userRepo.FindWishListByUid(id)
	if err != nil {
		return err
	}
	if wishlist.ID == 0 {
		err := w.userRepo.CreateWishList(id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *userUseCase) AddToWishListItem(uid, pfid uint) error {

	// get user wishist info
	wishlist, err := w.userRepo.FindWishListByUid(uid)
	if err != nil {
		return err
	}
	if wishlist.ID == 0 {
		res := errors.New("user does not have wishlist")
		return res
	}

	// check the product alredy exsit or not in the wishlist
	status, err := w.userRepo.FindProductFromWIshListItem(wishlist.ID, pfid)
	if err != nil {
		return err
	}
	if status {
		res := errors.New("product alredy exist")
		return res
	} else {
		err = w.userRepo.AddToWishlistItem(wishlist.ID, pfid)
		if err != nil {
			return err
		}
	}

	// response
	return nil
}

func (w *userUseCase) RemoveWishListItem(wid, pfid uint) error {

	wishlist, err := w.userRepo.FindWishListItemByWid(wid)
	if err != nil {
		if err != nil {
			return err
		}
	}

	if wishlist.Id != 0 {
		err := w.userRepo.RemoveFromWishListItem(wishlist.WishListId, pfid)
		if err != nil {
			return err
		}

	} else {
		res := errors.New("user does have wishlist item")
		return res
	}
	return nil
}

func (w *userUseCase) ViewWishList(uid uint, pagination req.PageNation) ([]res.ViewWishList, error) {

	var body []res.ViewWishList
	wishlist, err := w.userRepo.FindWishListByUid(uid)
	if err != nil {
		return body, err
	}

	if wishlist.ID == 0 {
		res := errors.New("user doee not have wishlist")
		return body, res
	}

	body, err = w.userRepo.ViewWishList(uid, pagination)
	if err != nil {
		return body, err
	}
	return body, nil
}
