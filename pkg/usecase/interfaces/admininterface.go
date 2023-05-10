package interfaces

import (
	"context"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
)

type AdminUseCase interface {
	Login(ctx context.Context, admin domain.Admin) (domain.Admin, error)
}
