package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/abhinandpn/project-ecom/pkg/config"
	"github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB: DB}
}

// Find admin via env loading
func (adm *adminDatabase) EnvAdminFind(ctx context.Context) (domain.Admin, error) {

	enfCfg := config.GetSudoAdminDetails()
	envAdmin := domain.Admin{
		Email:    enfCfg.AdminMail,
		Password: enfCfg.AdminPassword,
		Username: enfCfg.AdminUserName,
	}
	fmt.Println("email", envAdmin.Email, "pass", envAdmin.Password)
	if envAdmin.Email == "" || envAdmin.Password == "" || envAdmin.Username == "" {
		return envAdmin, errors.New("admin not found")
	}

	return envAdmin, nil
}

func (adm *adminDatabase) CreateAdmin(admin req.AdminLoginStruct) error {

	query := `Insert into admins (email,username,password)Values ($1,$2,$3)`

	if adm.DB.Exec(query, admin.Email, admin.UserName, admin.Password).Error != nil {
		return errors.New("faild to save admin to DB")
	}
	return nil
}

func (adm *adminDatabase) FindAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {

	query := `select * from admins where email=? or username=?`

	if adm.DB.Exec(query, admin.Email, admin.Username).Error != nil {
		return admin, errors.New("faild to find admin from DB")
	}
	return admin, nil
}

// List all users from database via pagenation
func (adm *adminDatabase) ListAllUser(ctx context.Context, PageNation req.PageNation) ([]res.UserResStruct, error) {

	var user []res.UserResStruct

	limit := PageNation.Count
	offset := (PageNation.PageNumber - 1) * limit

	query := `select * FROM users ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	err := adm.DB.Raw(query, limit, offset).Scan(&user).Error
	// fmt.Println("user data -------------->>> ", user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// block user
func (adm *adminDatabase) BlockUser(ctx context.Context, userId uint) error {

	// check the user id
	var user domain.Users
	query := `select * from users where id = ?`
	adm.DB.Exec(query, userId).Scan(&user)
	// validating the user
	/*
		validating for if we get any user id
		but the user does not exist in the database its will not
		provide any error
		its also give the empty spcae
		in the time the vaidation we can validate from the user variable
		the user is exist or not by cheking any users key like emil,number etc...!
	*/
	if user.Email == "" {
		return errors.New("invalid user id user does not exist")
	}

	// if we get the user
	// Start the function

	blockQry := `update users set block_status = $1 where id = $2`
	if adm.DB.Exec(blockQry, !user.IsBlocked, userId).Error != nil {
		return fmt.Errorf("faild to update block_status to %v", !user.IsBlocked)
	}
	return nil
}
