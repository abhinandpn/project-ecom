package interfaces

import (
	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
)

type Couponrepository interface {
	AddCoupon(coupon req.CouponReq) error
	UpdateCoupon(coupon req.CouponReq, Id uint) (domain.Coupon, error)
	DeleteCoupon(couponId uint) error
	ViewCoupons() ([]domain.Coupon, error)
	ViewCouponById(couponId uint) (domain.Coupon, error)
	ViewCouponByCode(code string)(domain.Coupon,error)
}
