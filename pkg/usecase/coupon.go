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
	cartRepo   interfaces.Cartrepository
}

func NewCouponUseCase(CouponRepo interfaces.Couponrepository,
	CartRepo interfaces.Cartrepository) service.CouponUseCase {
	return &CouponUseCase{couponRepo: CouponRepo,
		cartRepo: CartRepo}
}

func (cp *CouponUseCase) AddCoupon(coupon req.CouponWithMoney) error {

	body, err := cp.couponRepo.ViewCouponByCode(coupon.Code)
	if err != nil {
		return err
	}
	if body.Id != 0 {
		res := errors.New("code alredy exist")
		return res
	}
	err = cp.couponRepo.AddCouponWithMoney(coupon)
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

func (cp *CouponUseCase) UpdateCoupon(CouponId uint, coupon req.CouponWithMoney) (domain.Coupon, error) {

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

func (cp *CouponUseCase) ApplyCoupon(code string, uid uint) error {

	// find code
	coupon, err := cp.couponRepo.ViewCouponByCode(code)
	if err != nil {
		return err
	}
	if coupon.Id == 0 {
		res := errors.New("coupon does not exist")
		return res
	}
	// check if exist
	cart, err := cp.cartRepo.FindCartByUId(uid)
	if err != nil {
		return err
	}
	userCoupon, err := cp.couponRepo.FindCoupon(coupon.Id, uid)
	if err != nil {
		return err
	}
	if userCoupon.Id != 0 {
		res := errors.New("alredy have coupon")
		return res
	}
	if cart.CouponId == coupon.Id {
		res := errors.New("coupon alredy exist")
		return res
	}
	// find validity ---- updating

	// apply
	err = cp.couponRepo.ApplyCoupon(coupon.Id, uid)
	if err != nil {
		return err
	}
	return nil
}

func (cp *CouponUseCase) RemoveCoupon(code string, uid uint) error {

	coupon, err := cp.couponRepo.ViewCouponByCode(code)
	if err != nil {
		return err
	}

	cart, err := cp.cartRepo.FindCartByUId(uid)

	if err != nil {
		return err
	}
	if cart.CouponId != 0 {
		if coupon.Id == cart.CouponId {
			err = cp.couponRepo.RemoveCoupon(uid)
			if err != nil {
				return err
			}
		} else {
			res := errors.New("coupon does not exist  >>>>  1 ")
			return res
		}

	} else {
		res := errors.New("coupon does not exist >>> 2 ")
		return res
	}

	return nil
}
