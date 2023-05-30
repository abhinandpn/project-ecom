package res

// Full Cart response
type CartRes struct {
	Id       uint    `json:"id" gorm:"not null"`
	UserId   uint    `json:"user_id" gorm:"not null"`
	SubTotal float32 `json:"sub_total" gorm:"not null"`
	Total    float32 `json:"total" gorm:"not null"`
}

// List cart items
type CartItemsRes struct {
	Id          uint   `json:"id" gorm:"not null"`
	UserId      uint   `json:"user_id" gorm:"not null"`
	ProductId   uint   `json:"product_id" `
	ProductName string `json:"product_name"`
	Size        uint   `json:"size"`
	Category    string `json:"category"`
	Price       uint   `json:"price" gorm:"not null"`
}
