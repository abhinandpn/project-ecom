package req

type ReqProduct struct {
	ProductName string `json:"product_name" gorm:"not null" binding:"required,min=3,max=50"`
	Discription string `json:"description" gorm:"not null" binding:"required,min=3,max=100"`
	CategoryID  uint   `json:"category_id" binding:"required"`
	Price       uint   `json:"price" gorm:"not null" binding:"required,numeric"`
	Image       string `json:"image" gorm:"not null" binding:"required"`
	Color       string `json:"color" binding:"required"`
	Size        uint   `json:"size" binding:"required" gorm:"not null"`
	Brand       string `json:"brand" binding:"required"`
}
type PrDelReq struct {
	ProductName string `json:"product_name" gorm:"not null"`
}
type CategoryReq struct {
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
type UpdateCategoryReq struct {
	OldCategory string `json:"old_category" gorm:"not null"`
	Newcategory string `json:"new_category" gorm:""`
}
type SubCateCurdRes struct {
	// CategoryId      uint   `josn:"category_id"`
	SubCategoryName string `json:"sub_category_name"`
}
