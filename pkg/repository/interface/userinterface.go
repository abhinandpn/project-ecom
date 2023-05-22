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
	FindUserByNumber(ctx context.Context, number uint) (domain.Users, error)               // Find By Number
	FindUserById(ctx context.Context, id uint) (domain.Users, error)                       // Find By User Id
	FindUserByUserName(ctx context.Context, username string) (domain.Users, error)         // Find By UserName
	ListUsers(ctx context.Context, pagination req.PageNation) (res.ProductResponce, error) // List Full Users

	// opratios
	SaveUser(ctx context.Context, user domain.Users) (UserId uint, err error) // Create new User
	DeleteUser(ctx context.Context, id uint) error                            // Delete User
	UpdateUser(ctx context.Context, info domain.Users) (domain.Users, error)  // Update User Info
}
