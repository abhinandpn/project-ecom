package domain

import "time"

type Coupons struct {
	Id                    uint   `gorm:"primaryKey;unique;not null"`
	Code                  string ``
	DiscountPercent       float64
	DiscountMaximumAmount float64
	MinimumPurchaseAmount float64
	ExpirationDate        time.Time
}
