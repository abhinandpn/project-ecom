package domain

import "time"

type Product struct {
	Id            uint        `json:"id" gorm:"primaryKey;not null"`
	ProductName   string      `json:"product_name" gorm:"not null" binding:"required,min=3,max=50"`
	Discription   string      `json:"description" gorm:"not null" binding:"required,min=10,max=100"`
	CategoryID    uint        `json:"category_id" binding:"required omitempty,numeric" `
	Category      Category    `json:"-"`
	Brand         Brand       `json:"-"`
	Price         uint        `json:"price" gorm:"not null" binding:"required,numeric"`
	DiscountPrice uint        `json:"discount_price"`
	Info          ProductInfo `json:"-"`
	Image         string      `json:"image" gorm:"not null"`
	CreatedAt     time.Time   `json:"created_at" gorm:"not null"`
	UpdateAt      time.Time   `json:"updated_at"`
}
type ProductInfo struct {
	Id        uint   `json:"id" gorm:"primaryKey;not null"`
	ProductId uint   `json:"product_id" binding:"omitempty,numeric"`
	Colour    string `json:"colour"`
	Size      uint   `json:"size" binding:"required,numeric"`
	Brand     string `json:"brand"`
}

type Category struct {
	Id           uint   `json:"-" gorm:"primaryKey;not null"`
	CategoryName string `json:"category_name" gorm:"unique;not null" binding:"required,min=3,max=30"`
}
type Brand struct {
	Id        uint   `json:"-" gorm:"primaryKey;not null"`
	ProductId uint   `json:"product_id"`
	BrandName string `json:"brand_name" binding:"min=3,max=30"`
}
type ProductImage struct {
	Id            uint    `json:"id" gorm:"primarykey;not null"`
	ProductId     uint    `json:"product_id" gorm:"not null"`
	ProductImages string  `json:"product_images"`
	Product       Product `json:"-"`
}
