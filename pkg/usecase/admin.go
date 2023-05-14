package usecase

import (
	"context"
	"errors"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	services "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	adminRepo interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUseCase {

	return &AdminUseCase{adminRepo: repo}
}
func (adm *AdminUseCase) Login(ctx context.Context, admin domain.Admin) (domain.Admin, error) {
	// Get the admin from DB
	dbAdmin, err := adm.adminRepo.FindAdmin(ctx, admin)
	if err != nil {
		return admin, err
	} else if dbAdmin.ID == 0 {
		// return admin, errors.New("admin not exist")
		return admin, err

	}
	// Check the pass and username
	if bcrypt.CompareHashAndPassword([]byte(dbAdmin.Password), []byte(admin.Password)) != nil {
		return admin, errors.New("wrong password")
	}

	return dbAdmin, nil
}
