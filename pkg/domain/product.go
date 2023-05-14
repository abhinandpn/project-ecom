package domain

import "time"

type Product struct {
	Id            uint
	ProductName   string
	Discription   string
	CategoryId    uint
	Category      Category
	Price         uint
	DiscountPrice uint
	Image         string
	CreatedAt     time.Time
	UpdateAt      time.Time
}
type Category struct {
	Id           uint
	CategoryID   uint `json:"catrgory_id"`
	Category     *Category
	CategoryName string `json:"category_name" gorm:"unique;not null" binding:"required,min=3,max=30"`
}
