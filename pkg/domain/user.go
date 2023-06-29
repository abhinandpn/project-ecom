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
	// Default     *bool  `json:"is_default"`
}

type UserCart struct {
	Id     uint `json:"id" gorm:"not null"`
	UserId uint `json:"user_id" gorm:"not null"`
	User   Users
}

type CartInfo struct {
	Id            uint `json:"id" gorm:"not null"`
	CartId        uint `json:"cart_id" gorm:"not null"`
	ProductInfoId uint `json:"product_info_id"`
	Quantity      uint `json:"quantity"`
}

type WishList struct {
	ID     uint `json:"id" gorm:"primaryKey;not null"`
	UserID uint `json:"user_id" gorm:"not null"`
	User   Users
}
type WishListItems struct {
	Id            uint `json:"id"`
	WishListId    uint `json:"wish_list_id"`
	ProductInfoId uint `json:"product_info_id"`
}
