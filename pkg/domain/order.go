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
	CouponCode  uint
	TotalPrice  float64
	OrderStatus uint
	PaymentId   uint
	// Address     Address
}
type OrderItem struct {
	Id            uint `gorm:"primaryKey;unique;not null"`
	OrderInfo     uint
	ProductInfoId uint
	Quantity      uint
	// ProductInfo   ProductInfo
	// OrderInfo     OrderInfo
}

type OrderStatus struct {
	Id     uint `gorm:"primaryKey;unique;not null"`
	Status string
}
