package domain

import "time"

type Coupon struct {
	Id                 uint `gorm:"primaryKey;unique;not null"`
	Code               string
	DiscountPersentage string
	DiscountPrice      float64
	MinimumPurchase    float64
	ExpDate            time.Time
}
