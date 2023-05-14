package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/abhinandpn/project-ecom/pkg/config"
	domain "github.com/abhinandpn/project-ecom/pkg/domain"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)

	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	err := db.AutoMigrate(

		// User Table
		&domain.Users{},
		&domain.UserInfo{},
		&domain.Address{},

		// Admin Table
		&domain.Admin{},

		// Product Table
		&domain.Category{},
		&domain.ProductInfo{},
		&domain.Product{},
		&domain.ProductImage{},
		&domain.Brand{},

		// &domain.Coupons{},

	)

	// Error handling While Database Table creating
	if err != nil {
		return nil, err
	}

	return db, dbErr
}
