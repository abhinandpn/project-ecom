package interfaces

import (
	"context"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
)

type UserUseCase interface {

	// Signup
	SignUp(ctx context.Context, user domain.Users) error
	Login(ctx context.Context, user domain.Users) (domain.Users, error)
	OtpLogin(ctx context.Context, user domain.Users) (domain.Users, error)

	UserAccount(ctx context.Context, UserId uint) (domain.Users, error)
	//........................................................
	/*
		Add Address
		Uodate Addres
		List All Address
		Delete Drress
		View profile
		Edit Profile
	*/
}
