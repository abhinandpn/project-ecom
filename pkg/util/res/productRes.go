package res

import "time"

type ProductResponce struct {
	ID            uint      `json:"product_id"`
	ProductName   string    `json:"product_name"`
	Discription   string    `json:"description" `
	CategoryName  string    `json:"category_name"`
	Size          uint      `json:"size"`
	Price         uint      `json:"price"`
	DiscountPrice uint      `json:"discount_price"`
	Image         string    `json:"image"`
	CreatedAt     time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt     time.Time `json:"updated_at"`
	// CategoryID    uint      `json:"category_id"`
}
type CategoryRes struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
}
type SubCategoryRes struct {
	Id              uint   `json:"id"`
	CategoryId      uint   `json:"category_id"`
	CategoryName    string `json:"category_name"`
	SubcategoryName string `json:"sub_category_name"`
}
type BrandRes struct {
	Id        uint   `json:"id"`
	BrandName string `json:"brand_name"`
}
type ProductQuentity struct {
	ProductName  string
	CategoryName string
	BrandName    string
	Price        float64
	Colour       string
	Size         uint
	ProductImage string
	Quentity     uint
}
type ResProduct struct {
	ProductName string `json:"product_name" gorm:"not null" binding:"required,min=3,max=50"`
	Discription string `json:"description" gorm:"not null" binding:"required,min=3,max=100"`
	CategoryID  uint   `json:"category_id" binding:"required"`
	BrandId     uint   `json:"brand_id" binding:"required"`
	Price       uint   `json:"price" gorm:"not null" binding:"required,numeric"`
	Color       string `json:"color" binding:"required"`
	Size        uint   `json:"size" binding:"required" gorm:"not null"`
	Image       string `json:"image" gorm:"not null" binding:"required"`
}

type ResBrand struct {
	BrandName  string `json:"brand_name"`
	BrandImage string `json:"brand_image"`
}

/*
SELECT

	p.product_name,
	sc.sub_category_name,
	br.brand_name,
	pi.price,
	pi.colour,
	pi.size,
	pr.product_images,
	pi.quatity

FROM

	products p
	JOIN product_infos pi ON p.id = pi.product_id
	JOIN brands br ON p.brand_id = br.id
	JOIN product_images pr ON p.id = pr.product_id
	JOIN sub_categories sc ON p.sub_category_id = sc.category_id;
*/
type ResProductOrder struct {
	ProductName  string
	Discription  string
	CategoryName string
	BrandName    string
	Size         uint
	Price        float64
	Colour       string
}
