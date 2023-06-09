package repository

import (
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"gorm.io/gorm"
)

type cartDatabase struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.Cartrepository {

	return &cartDatabase{DB: db}
}
