package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
)

type AdminRepository interface {
	CreateAdmin(req.AdminLoginStruct) error
	FindAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error)
}
