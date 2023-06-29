package domain

import "time"

type UserOrder struct {
	Id     uint `gorm:"primaryKey;unique;not null"`
	UserId uint
	// Users  Users
}
type OrderInfo struct {
	Id          uint `gorm:"primaryKey;unique;not null"`
	OrderId     uint
	OrderTime   time.Time
	AddressId   uint
	CouponCode  string
	TotalPrice  float64
	OrderStatus string
	PaymentId   string
	// Address     Address
}
type OrderItem struct {
	Id            uint `gorm:"primaryKey;unique;not null"`
	UserOrderId   uint
	ProductInfoId uint
	Quantity      uint
	UserCartId    uint
	// ProductInfo   ProductInfo
	// OrderInfo     OrderInfo
}
