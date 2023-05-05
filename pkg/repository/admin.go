package repository

import (
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"gorm.io/gorm"
)

type AdminDB struct {
	DB *gorm.DB
}

func NewAdminrepo(DB *gorm.DB) interfaces.AdminRepository {
	return &AdminDB{DB: DB}
}
