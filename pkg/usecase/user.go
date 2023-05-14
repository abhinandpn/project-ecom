package usecase

import (
	"context"
	"errors"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/helper"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
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
			// return errors.New("error to hash the password")
			return err

		}

		// save the password in hash
		user.Password = string(hashpass)

		// save the user
		_, err = usr.userRepo.SaveUser(ctx, user)
		if err != nil {
			// return errors.New("error to save user")
			return err

		}
		return nil
	}

	// if user exist then check which field is exist

	return helper.UserCheck(user, checkUser)

}

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
