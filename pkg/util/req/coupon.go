package req

import "time"

type CouponReq struct {
	Code               string
	DiscountPersentage string
	DiscountPrice      float64
	MinimumPurchase    float64
	ExpDate            time.Time
}

type CouponWithMoney struct {
	Code            string
	DiscountPrice   float64
	MinimumPurchase float64
	ExpDate         time.Time
}

type CouponWithPercentage struct {
	Code            string
	DiscountPrice   float64
	MinimumPurchase float64
	ExpDate         time.Time
}
