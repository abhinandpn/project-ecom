package res

import "time"

type ProductResponce struct {
	ID            uint      `json:"product_id"`
	ProductName   string    `json:"product_name"`
	Description   string    `json:"description" `
	CategoryID    uint      `json:"category_id"`
	CategoryName  string    `json:"category_name"`
	Size          uint      `json:"size"`
	Price         uint      `json:"price"`
	DiscountPrice uint      `json:"discount_price"`
	Image         string    `json:"image"`
	CreatedAt     time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt     time.Time `json:"updated_at"`
}
type CategoryRes struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
}
type SubCategoryRes struct {
	Id              uint   `json:"id"`
	SubcategoryName string `json:"sub_category_name"`
	CategoryId      uint   `json:"category_id"`
}
type BrandRes struct {
	Id        uint   `json:"id"`
	BrandName string `json:"brand_name"`
}
