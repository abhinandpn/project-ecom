package repository

import (
	"context"
	"errors"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}
func (adm *adminDatabase) FindAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {
	if adm.DB.Raw("SELECT * FROM admins WHERE email=? OR username=?", admin.Email, admin.Username).Scan(&admin).Error != nil {
		return admin, errors.New("faild to find admin")
	}
	return admin, nil
}
