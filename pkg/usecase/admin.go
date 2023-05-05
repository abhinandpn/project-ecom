package usecase

import (
	"context"
	"errors"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	services "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepo interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepo: repo,
	}
}
func (c *adminUseCase) SignUp(ctx context.Context, admin domain.Admin) error {

	// find if the admin exist or not
	if admin, err := c.adminRepo.FindAdmin(ctx, admin); err != nil {
		return err
	} else if admin.ID != 0 {
		return errors.New("can't save admin already exist with this details")
	}
	// if its new a entry creating admin

	// generating hash pass for admin
	hashPass, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return errors.New("faild to generate hashed password for admin")
	}
	// setting hashed pass 	on admin
	admin.Password = string(hashPass)

	return c.adminRepo.SaveAdmin(ctx, admin)
}
