package interfaces

import (
	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
)

type CouponUseCase interface {
	AddCoupon(coupon req.CouponReq) error
	DeleteCoupon(couponId uint) error
	ViewCouponById(couponId uint) (domain.Coupon, error)
	ViewCouponByCode(name string) (domain.Coupon, error)
	ListCoupon() ([]domain.Coupon, error)
	UpdateCoupon(CouponId uint, coupon req.CouponReq) (domain.Coupon, error)
}
