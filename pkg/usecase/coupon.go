package usecase

import (
	"errors"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
)

type CouponUseCase struct {
	couponRepo interfaces.Couponrepository
}

func NewCouponUseCase(CouponRepo interfaces.Couponrepository) service.CouponUseCase {
	return &CouponUseCase{couponRepo: CouponRepo}
}

func (cp *CouponUseCase) AddCoupon(coupon req.CouponReq) error {

	body, err := cp.couponRepo.ViewCouponByCode(coupon.Code)
	if err != nil {
		return err
	}
	if body.Id != 0 {
		res := errors.New("code alredy exist")
		return res
	}
	err = cp.AddCoupon(coupon)
	if err != nil {
		return err
	}
	return nil
}

func (cp *CouponUseCase) DeleteCoupon(couponId uint) error {

	body, err := cp.couponRepo.ViewCouponById(couponId)
	if err != nil {
		return err
	}
	if body.Id == 0 {
		res := errors.New("code does not exist")
		return res
	}
	err = cp.couponRepo.DeleteCoupon(couponId)
	if err != nil {
		return err
	}
	return nil
}

func (cp *CouponUseCase) ViewCouponById(couponId uint) (domain.Coupon, error) {

	var body domain.Coupon
	body, err := cp.couponRepo.ViewCouponById(couponId)
	if err != nil {
		return body, err
	}
	if body.Id == 0 {
		res := errors.New("code does not exist")
		return body, res
	}
	body, err = cp.couponRepo.ViewCouponById(couponId)
	if err != nil {
		return body, err
	}
	return body, nil
}

func (cp *CouponUseCase) UpdateCoupon(CouponId uint, coupon req.CouponReq) (domain.Coupon, error) {

	body, err := cp.couponRepo.ViewCouponById(CouponId)
	if err != nil {
		return body, err
	}
	if body.Id != 0 {
		body, err = cp.couponRepo.UpdateCoupon(coupon, CouponId)
		if err != nil {
			return body, err
		}
	} else {
		res := errors.New("code does not exist")
		return body, res
	}
	return body, nil
}

func (cp *CouponUseCase) ListCoupon() ([]domain.Coupon, error) {

	var body []domain.Coupon
	body, err := cp.couponRepo.ViewCoupons()
	if err != nil {
		return body, err
	}
	if body == nil {
		res := errors.New("coupon does  not exist")
		return body, res
	}
	return body, nil
}

func (cp *CouponUseCase) ViewCouponByCode(name string) (domain.Coupon, error) {

	body, err := cp.couponRepo.ViewCouponByCode(name)
	if err != nil {
		return body, err
	}
	if body.Id == 0 {
		res := errors.New("coupon code does not exist")
		return body, res
	}
	return body, nil
}
