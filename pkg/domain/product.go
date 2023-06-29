package domain

import "time"

type Category struct {
	Id           uint   `json:"id" gorm:"primaryKey;not null"`
	CategoryName string `json:"category_name" gorm:"unique;not null" binding:"required,min=3,max=30"`
}
type SubCategory struct {
	Id              uint   `json:"id" gorm:"primaryKey;not null"`
	CategoryId      uint   `json:"category_id" binding:"omitempty,numeric"`
	SubCategoryName string `json:"sub_category_name"`
}

// New Updated ProductInfo

type Product struct {
	Id            uint      `json:"id" gorm:"primaryKey;not null"`
	ProductName   string    `json:"product_name" gorm:"not null" binding:"required,min=3,max=50"`
	Discription   string    `json:"discription" gorm:"not null" binding:"required,min=3,max=50"`
	BrandId       uint      `json:"brand_id" gorm:"not null"`
	CategoryId    uint      `json:"category_id" gorm:"not null"`
	SubCategoryId uint      `json:"sub_catgory_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Info          ProductInfo
}
type ProductInfo struct {
	Id        uint    `json:"id" gorm:"primaryKey;not null"`
	ProductId uint    `json:"product_id" gorm:"not null"`
	Price     float64 `json:"price" gorm:"not null"`
	Colour    string  `json:"colour"`
	Size      uint    `json:"size" gorm:"not null"`
	// ImageId   uint    `json:"image_id" gorm:"not null"`
	Quatity uint `json:"quantity"`
}
type Brand struct {
	Id         uint   `json:"id" gorm:"primaryKey;not null"`
	BrandName  string `json:"brand_name" binding:"min=3,max=30"`
	BrandImage string `json:"brand_image"`
}

type ProductImage struct {
	Id            uint    `json:"id" gorm:"primarykey;not null"`
	ProductId     uint    `json:"product_id" gorm:"not null"`
	ProductImages string  `json:"product_images"`
	Product       Product `json:"-"`
}
