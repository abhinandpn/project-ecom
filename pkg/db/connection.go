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
	// User Table
	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.UserInfo{})
	db.AutoMigrate(&domain.Address{})

	// Admin Table
	db.AutoMigrate(&domain.Admins{})

	// Product Table
	// cart Table
	// Wish List
	// order Table
	// Payment Detail

	return db, dbErr
}
