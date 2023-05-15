package domain

import "time"

type Product struct {
	Id            uint        `json:"id" gorm:"primaryKey;not null"`
	ProductName   string      `json:"product_name" gorm:"not null" binding:"required,min=3,max=50"`
	Discription   string      `json:"description" gorm:"not null" binding:"required,min=10,max=100"`
	CategoryId    uint        `json:"category_id" binding:"omitempty,numeric"`
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
	Brand     uint   `json:"brand"`
}

/*
in Category : Id means the main ctegoy id
in category : CategoryId means subcategory
*/
type Category struct {
	Id           uint      `json:"-" gorm:"primaryKey;not null"`
	CategoryID   uint      `json:"catrgory_id"`
	Category     *Category `json:"-"`
	CategoryName string    `json:"category_name" gorm:"unique;not null" binding:"required,min=3,max=30"`
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

// Add On Tables
/*
	Brand table
	Colour Table
	Size table

	Updation --> Sub Category Table
*/
