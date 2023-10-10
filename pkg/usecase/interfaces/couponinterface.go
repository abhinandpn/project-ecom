package interfaces

import (
	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
)

type CouponUseCase interface {
	AddCoupon(coupon req.CouponWithMoney) error
	DeleteCoupon(couponId uint) error
	ViewCouponById(couponId uint) (domain.Coupon, error)
	ViewCouponByCode(name string) (domain.Coupon, error)
	ListCoupon() ([]domain.Coupon, error)
	UpdateCoupon(CouponId uint, coupon req.CouponWithMoney) (domain.Coupon, error)

	//
	ApplyCoupon(code string, uid uint) error
	RemoveCoupon(code string, uid uint) error
	GetAppliedCoupon(uid uint) (domain.Coupon, error)
}
