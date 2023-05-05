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
		&domain.Categery{},
		&domain.ProductBrand{},
		&domain.Product{},
		&domain.ProductItem{},
		&domain.Coupons{},

		// cart Table
		&domain.Cart{},
		&domain.CartItems{},

		// Wish List
		&domain.Wishlist{},
		&domain.WishlistItem{},

		// order Table
		&domain.Order{},
		&domain.OrderLine{},
		&domain.OrderStatus{},
		&domain.DeliveryStatus{},
		&domain.Return{},

		// Payment Detail
		&domain.PaymentStatus{},
		&domain.PaymentDetails{},
	)

	// Error handling For Database Table creating
	if err != nil {
		return nil, err
	}

	return db, dbErr
}
