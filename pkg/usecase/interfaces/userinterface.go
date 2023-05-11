package interfaces

import (
	"context"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
)

type UserUseCase interface {

	// Signup
	SignUp(ctx context.Context, user domain.Users) error

	//........................................................
	FindAll(ctx context.Context) ([]domain.Users, error)
	FindByID(ctx context.Context, id uint) (domain.Users, error)
	Save(ctx context.Context, user domain.Users) (domain.Users, error)
	Delete(ctx context.Context, user domain.Users) error
}
