package interfaces

import (
	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
)

type Couponrepository interface {
	AddCouponWithMoney(coupon req.CouponWithMoney) error
	UpdateCoupon(coupon req.CouponWithMoney, Id uint) (domain.Coupon, error)
	DeleteCoupon(couponId uint) error
	ViewCoupons() ([]domain.Coupon, error)
	ViewCouponById(couponId uint) (domain.Coupon, error)
	ViewCouponByCode(code string) (domain.Coupon, error)

	//
	ApplyCoupon(cid, uid uint) error // cart apply
	RemoveCoupon(uid uint) error
	FindCoupon(cid, uid uint) (domain.Coupon, error)
}
