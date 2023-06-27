package usecase

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/helper"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) service.UserUseCase {
	return &userUseCase{userRepo: repo}
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
		fmt.Println("--------------", Uid)
		err = usr.userRepo.CreateWishList(Uid)
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

	err := usr.userRepo.AddAddress(ctx, Uid, Address)

	return err

}

func (usr *userUseCase) UpdateAddress(ctx context.Context, Uid uint, address req.ReqAddress) error {

	var body req.ReqAddress

	err := usr.userRepo.UpdateAddress(ctx, Uid, body)

	return err
}
func (usr *userUseCase) ListAllAddress(ctx context.Context, Uid uint) ([]res.ResAddress, error) {

	var body []res.ResAddress

	body, err := usr.userRepo.ListAllAddress(ctx, Uid)

	return body, err
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
	fmt.Println("wish list id (usecase 187)", wishlist.WishListId)
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

func (w *userUseCase) AddToWishListItem(wid, pfid uint) error {

	wishlist, err := w.userRepo.FindWishListItemByWid(wid)
	if err != nil {
		if err != nil {
			return err
		}
	}

	if wishlist.Id == 0 {
		body, err := w.userRepo.FindProductFromWIshListItem(wishlist.Id, pfid)
		if err != nil {
			return err
		}
		if !body {
			err := w.userRepo.AddToWishlistItem(wid, pfid)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (w *userUseCase) RemoveWishListItem(wid, pfid uint) error {

	wishlist, err := w.userRepo.FindWishListItemByWid(wid)
	if err != nil {
		if err != nil {
			return err
		}
	}
	fmt.Println("wishlis (247 usecase)", wishlist)
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
