package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type AdminRepository interface {
	// Login Admin sudo .env
	EnvAdminFind(ctx context.Context) (domain.Admin, error)
	// Admin Main
	CreateAdmin(req.AdminLoginStruct) error
	FindAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error)
	// user Side
	ListAllUser(ctx context.Context, PageNation req.PageNation) (user []res.UserResStruct, err error)
	BlockUser(ctx context.Context, userId uint) error
}
