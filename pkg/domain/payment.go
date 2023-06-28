package domain

import "time"

type PaymentMethod struct {
	Id     uint
	Method string
}

type PaymentDetail struct {
	Id              uint
	OrderId         uint
	TotalPrice      float64
	PaymentMethodId uint
	PaymentStatusId uint
	PaymentRef      string
	UpdatedAt       time.Time
}

type PaymentStatus struct {
	Id            uint
	PaymentStatus string
}
