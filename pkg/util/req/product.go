package req

type ReqProduct struct {
	ProductName string `json:"product_name" gorm:"not null" binding:"required,min=3,max=50"`
	Discription string `json:"description" gorm:"not null" binding:"required,min=3,max=100"`
	CategoryID  uint   `json:"category_id" binding:"required"`
	BrandId     uint   `json:"brand_id" binding:"required"`
	Price       uint   `json:"price" gorm:"not null" binding:"required,numeric"`
	Color       string `json:"color" binding:"required"`
	Size        uint   `json:"size" binding:"required" gorm:"not null"`
	// ImageId     uint   `json:"image_id"`
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
	BrandName  string `json:"brand_name"`
	BrandImage string `json:"brand_image"`
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

//  ---------- Sorting Struct ----------

type SortReqColour struct {
	Colour string `json:"colour"`
}
type SortReqSize struct {
	Size int `json:"size"`
}
type SortReqCategory struct {
	Category string `json:"category"`
}
type SortReqBrand struct {
	Brand string `json:"brand"`
}
type SortReqName struct {
	Name string `json:"name"`
}
type SortReqPrice struct {
	PriceStart int `json:"price_start"`
	PriceEnd   int `json:"price_end"`
}
type SortReqQuantity struct {
	QuantityStart int `json:"quantity_start"`
	QuantityEnd   int `json:"quantity_end"`
}

// ------- sorting end -------
type UpdateProduct struct {
	ProductName string
	Discription string
	BrandId     uint
	CategoryId  uint
	Price       float64
	Colour      string
	Size        uint
	Quantity    uint
	// Image       string
}
