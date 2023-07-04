package domain

import "time"

type PaymentMethod struct {
	Id     uint `json:"id" gorm:"primaryKey;not null"`
	Method string
}

type PaymentDetail struct {
	Id              uint `json:"id" gorm:"primaryKey;not null"`
	OrderId         uint
	TotalPrice      float64
	PaymentMethodId uint
	PaymentStatusId uint
	PaymentRef      string
	UpdatedAt       time.Time
}

type PaymentStatus struct {
	Id            uint `json:"id" gorm:"primaryKey;not null"`
	PaymentStatus string
}
