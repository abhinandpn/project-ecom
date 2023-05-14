package repository

import (
	"context"
	"errors"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB: DB}
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
