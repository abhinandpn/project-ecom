package usecase

import (
	"context"
	"errors"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/helper"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	services "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
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
			return errors.New("error to hash the password")

		}

		// save the password in hash
		user.Password = string(hashpass)

		// save the user
		_, err = usr.userRepo.SaveUser(ctx, user)
		if err != nil {
			return errors.New("error to save user")

		}
		return nil
	}

	// if user exist then check which field is exist

	return helper.UserCheck(user, checkUser)

}

// ........................................
func (c *userUseCase) FindAll(ctx context.Context) ([]domain.Users, error) {
	users, err := c.userRepo.FindAll(ctx)
	return users, err
}

func (c *userUseCase) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	user, err := c.userRepo.FindByID(ctx, id)
	return user, err
}

func (c *userUseCase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	user, err := c.userRepo.Save(ctx, user)

	return user, err
}

func (c *userUseCase) Delete(ctx context.Context, user domain.Users) error {
	err := c.userRepo.Delete(ctx, user)

	return err
}
