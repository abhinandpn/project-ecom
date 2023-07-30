package req

import "time"

type CouponReq struct {
	Code               string
	DiscountPersentage string
	DiscountPrice      float64
	MinimumPurchase    float64
	ExpDate            time.Time
}
