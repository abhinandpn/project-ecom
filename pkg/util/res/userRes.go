package res

import "time"

type UserResStruct struct {
	ID          uint      `json:"id" copier:"must"`
	FName       string    `json:"first_name" copier:"must"`
	LName       string    `json:"last_name" copier:"must"`
	Age         uint      `json:"age" copier:"must"`
	Email       string    `json:"email" copier:"must"`
	UserName    string    `json:"user_name" copire:"must"`
	Number      string    `json:"phone" copier:"must"`
	BlockStatus bool      `json:"block_status" copier:"must"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type UserFindWithUserName struct {
	UserName string `json:"user_name" gorm:"not null"`
}
type ResAddress struct {
	Id          uint   `json:"id"`
	UserId      uint   `json:"user_id" copier:"must"`
	Name        string `json:"name" copier:"must"`
	House       string `json:"house" copier:"must"`
	PhoneNumber string `json:"phone_number" copier:"must"`
	Street      string `json:"street" copier:"must"`
	City        string `json:"city" copier:"must"`
	District    string `json:"district" copier:"must"`
	Pincode     string `json:"pincode" copier:"must"`
	Landmark    string `json:"landmark" copier:"must"`
	// CountryID   uint   `json:"country_id" binding:"required"`

	Default *bool
}

type ViewWishList struct {
	Id          uint
	ProductName string
	Price       float64
	Colour      string
	Brand       string
	Category    string
}
type UserBlockRes struct {
	UserName    string
	Email       string
	Number      string
	BlockStatus bool
}
