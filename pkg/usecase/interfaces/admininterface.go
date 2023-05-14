package interfaces

import (
	"context"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type AdminUseCase interface {
	// Admin side
	Login(ctx context.Context, admin domain.Admin) (domain.Admin, error)
	// User side
	FindAllUser(ctx context.Context, pagination req.PageNation) (users []res.UserResStruct, err error)
	BlockUser(ctx context.Context, UserId uint) error
}
