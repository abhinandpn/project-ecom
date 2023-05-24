package interfaces

import (
	"context"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type AdminUseCase interface {

	// Sudo admin .env file loading
	SudoLogin(ctx context.Context, admin domain.Admin) (domain.Admin, error)

	// Admin side
	Login(ctx context.Context, admin domain.Admin) (domain.Admin, error)

	// -------------User side----------------

	// List All Users
	FindAllUser(ctx context.Context, pagination req.PageNation) (users []res.UserResStruct, err error)
	BlockUser(ctx context.Context, UserId uint) error

	// Find user By
	FindUserByUserName(ctx context.Context, name string) (domain.Users, error)
	FindUserByNumber(ctx context.Context, number string) (domain.Users, error)
	FindUserByEmail(ctx context.Context, email string) (domain.Users, error)
	FindUserInfo(ctx context.Context, user domain.Users) (domain.Users, error)
	FindUserById(ctx context.Context, id uint) (domain.Users, error)
}
