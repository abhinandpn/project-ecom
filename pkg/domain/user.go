package domain

import "time"

type Users struct {
	ID        uint   `json:"id" gorm:"unique;not null"`
	UserName  string `json:"username" binding:"required"`
	FName     string `json:"f_name"`
	LName     string `json:"l_name"`
	Email     string `gorm:"uniqueIndex" json:"email" binding:"required" validate:"required,email"`
	Number    string `gorm:"uniqueIndex" json:"phone" validate:"required,number"`
	Password  string `json:"password" binding:"required" validate:"required,min=8,max=64"`
	IsBlocked bool   `json:"isblock" gorm:"not null;default:false"`
	CreatedAt time.Time
}
type UserInfo struct {
	ID                uint `gorm:"primaryKey"`
	UsersID           uint
	Users             Users `gorm:"foreignKey:UsersID"`
	BlockedAt         time.Time
	BlockedBy         uint
	ReasonForBlocking string
}

type Address struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Users       Users  `gorm:"foreignKey:UserID" json:"-"`
	House       string `json:"house8"`
	PhoneNumber string `json:"phone_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	Pincode     string `json:"pincode"`
	Landmark    string `json:"landmark"`
	Default     *bool  `json:"is_default"`
}

type Cart struct {
	Id              uint    `json:"id" gorm:"primaryKey;not null"`
	UserID          uint    `json:"user_id" gorm:"not null"`
	AppliedCouponID uint    `json:"applied_coupon_id"`
	DiscountAmount  float64 `json:"discount_amount"`
	TotalPrice      float64 `json:"total_price" gorm:"not null"`
}

type CartIteams struct {
	CartItemId uint `json:"cart_item_id" gorm:"not null"`
	CartId     uint `json:"cart_id" gorm:"not null"`
	Cart       Cart
	ProductId  uint `json:"product_id" gorm:"not null"`
	Product    Product
	Quantity   uint `json:"qty" gorm:"not null"`
}

type WishList struct {
	ID        uint `json:"id" gorm:"primaryKey;not null"`
	UserID    uint `json:"user_id" gorm:"not null"`
	User      Users
	ProductId uint `json:"product_id" gorm:"not null"`
}
