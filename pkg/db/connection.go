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
		&domain.WishList{},
		&domain.WishListItems{},
		&domain.UserOrder{},
		&domain.OrderInfo{},
		&domain.OrderItem{},
		&domain.OrderStatus{},
		&domain.PaymentMethod{},
		&domain.PaymentDetail{},
		&domain.PaymentStatus{},

		// Admin Table
		&domain.Admin{},

		// Product Table
		&domain.Category{},
		&domain.Coupon{},

		// product updated {product branch}
		&domain.Product{},
		&domain.ProductInfo{},
		&domain.ProductImage{},
		&domain.Brand{},

		// Cart Table
		&domain.UserCart{},
		&domain.CartInfo{},
		// &domain.WishList{},
		&domain.SubCategory{},

		// &domain.Coupons{},

		// sub category updating

	)

	// Error handling While Database Table creating
	if err != nil {
		return nil, err
	}

	return db, dbErr
}
