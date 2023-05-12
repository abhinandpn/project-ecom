package domain

import "time"

type Users struct {
	ID        uint   `json:"id" gorm:"unique;not null"`
	UserName  string `json:"username" binding:"required"`
	FName     string `json:"f_name"`
	LName     string `json:"l_name"`
	Email     string `gorm:"uniqueIndex" json:"email" binding:"required" validate:"required,email"`
	Number    int    `gorm:"uniqueIndex" json:"phone" validate:"required,number"`
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
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	Pincode     string `json:"pincode"`
	Landmark    string `json:"landmark"`
}
