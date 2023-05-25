package req

type ReqProduct struct {
	ProductName string `json:"product_name" gorm:"not null" binding:"required,min=3,max=50"`
	Discription string `json:"description" gorm:"not null" binding:"required,min=10,max=100"`
	CategoryID  uint   `json:"category_id" binding:"required"`
	Size        uint   `json:"size" binding:"required"`
	Color       string `json:"color" binding:"required"`
	Brand       string `json:"brand" binding:"required"`
	Price       uint   `json:"price" gorm:"not null" binding:"required,numeric"`
	Image       string `json:"image" gorm:"not null" binding:"required"`
}
type CategoryReq struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
}
type SubCategoryReq struct {
	Id              uint   `json:"id"`
	SubcategoryName string `json:"sub_category_name"`
	CategoryId      uint   `json:"category_id"`
}
type BrandReq struct {
	Id        uint   `json:"id"`
	BrandName string `json:"brand_name"`
}
type AddCategoryReq struct {
	CategoryName string `json:"category_name" gorm:"not null"`
}
