package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
)

type adminUseCase interface {
	// Admin interface
	SignUp(ctx context.Context, admin domain.Admin) error
}
