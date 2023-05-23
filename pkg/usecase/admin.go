package usecase

import (
	"context"
	"errors"
	"fmt"

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
	userRepo  interfaces.UserRepository
}

func NewAdminUseCase(Adminrepo interfaces.AdminRepository, UserRepo interfaces.UserRepository) services.AdminUseCase {

	return &AdminUseCase{
		adminRepo: Adminrepo,
		userRepo:  UserRepo,
	}
}

// Sudo admin login .env file loading
func (adm *AdminUseCase) SudoLogin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {

	envAdmin, err := adm.adminRepo.EnvAdminFind(ctx)
	fmt.Println(envAdmin.Email, admin.Email)
	if err != nil {
		return domain.Admin{}, err
	}

	if envAdmin.Email == admin.Email && envAdmin.Password == admin.Password {
		return envAdmin, nil
	} else if envAdmin.Username == admin.Username && envAdmin.Password == admin.Password {
		return envAdmin, nil
	} else {
		return admin, errors.New("invalid admin")
	}
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

// Find User By user Name
func (adm *AdminUseCase) FindUserByUserName(ctx context.Context, name string) (domain.Users, error) {

	body, err := adm.userRepo.FindUserByUserName(ctx, name)

	if err != nil {
		return body, err
	}

	return body, nil
}

// Find User By Email
func (adm *AdminUseCase) FindUserByEmail(ctx context.Context, email string) (domain.Users, error) {

	body, err := adm.userRepo.FindUserByEmail(ctx, email)

	if err != nil {
		return body, err
	}
	return body, nil
}

// Find User By Number
func (adm *AdminUseCase) FindUserByNumber(ctx context.Context, number uint) (domain.Users, error) {

	body, err := adm.userRepo.FindUserByNumber(ctx, number)
	if err != nil {
		return body, err
	}
	return body, nil
}

// Find User By Any Information

func (adm *AdminUseCase) FindUserInfo(ctx context.Context, user domain.Users) (domain.Users, error) {

	body, err := adm.userRepo.FindUser(ctx, user)

	if err != nil {
		return body, err
	}

	return body, nil

}

// Find User By Id
func (adm *AdminUseCase) FindUserById(ctx context.Context, id uint) (domain.Users, error) {

	body, err := adm.userRepo.FindUserById(ctx, id)

	if err != nil {
		return body, err
	}

	return body, nil
}
