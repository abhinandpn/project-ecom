package req

type CartReq struct {
	ProductId       uint
	AppliedCouponID string
	DiscountAmount  float64
}
