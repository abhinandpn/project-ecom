package usecase

import (
	"context"
	"errors"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	services "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"github.com/jinzhu/copier"
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

// Find all user
func (adm *AdminUseCase) FindAllUser(ctx context.Context, pagination req.PageNation) (users []res.UserResStruct, err error) {

	users, err = adm.adminRepo.ListAllUser(ctx, pagination)

	if err != nil {
		return nil, err
	}

	var responce []res.UserResStruct
	copier.Copy(&responce, &users)
	return responce, nil
}

// Block user
func (adm *AdminUseCase) BlockUser(ctx context.Context, UserId uint) error {

	return adm.adminRepo.BlockUser(ctx, UserId)
}
