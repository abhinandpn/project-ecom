package interfaces

import (
	"context"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
)

type UserUseCase interface {

	// Signup
	SignUp(ctx context.Context, user domain.Users) error

	//........................................................
}
