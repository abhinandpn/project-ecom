package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
)

type UserRepository interface {

	// ..............................................
	FindUser(ctx context.Context, user domain.Users) (domain.Users, error)
	FindUserByEmail(ctx context.Context, email string) (user domain.Users, err error)
	FindUserByNumber(ctx context.Context, number uint) (user domain.Users, err error)
	SaveUser(ctx context.Context, user domain.Users) (UserId uint, err error)
	// ..............................................

	FindAll(ctx context.Context) ([]domain.Users, error)
	FindByID(ctx context.Context, id uint) (domain.Users, error)
	Save(ctx context.Context, user domain.Users) (domain.Users, error)
	Delete(ctx context.Context, user domain.Users) error
}
