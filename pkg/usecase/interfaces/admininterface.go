package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
)

type AdminUseCase interface {
	// Admin interface
	Login(ctx context.Context, admin domain.Admin) error
}
