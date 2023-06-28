package domain

import "time"

type UserOrder struct {
	Id     uint
	UserId uint
	Users  Users
}
type OrderInfo struct {
	Id          uint
	UserOrderId uint
	OrderTime   time.Time
	AddressId   uint
	Address     Address
	CouponCode  string
	TotalPrice  float64
	OrderStatus string
}
type OrderItem struct {
	Id            uint
	UserOrderId   uint
	OrderInfo     OrderInfo
	ProductInfoId uint
	ProductInfo   ProductInfo
	Quantity      uint
}
