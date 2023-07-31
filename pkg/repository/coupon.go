package repository

import (
	"github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"gorm.io/gorm"
)

type CouponDatabase struct {
	DB *gorm.DB
}

func NewCouponRepository(db *gorm.DB) interfaces.Couponrepository {

	return &CouponDatabase{DB: db}
}

func (cp *CouponDatabase) AddCouponWithMoney(coupon req.CouponWithMoney) error {

	var body domain.Coupon
	query := `insert into coupons (code,
					discount_price,
					minimum_purchase,
					exp_date)values ($1,$2,$3,$4);`
	err := cp.DB.Raw(query,
		coupon.Code,
		coupon.DiscountPrice,
		coupon.MinimumPurchase,
		coupon.ExpDate).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (cp *CouponDatabase) UpdateCoupon(coupon req.CouponWithMoney, Id uint) (domain.Coupon, error) {

	var body domain.Coupon
	query := `update coupons set code = $1,
						discount_price = $2,
						minimum_purchase = $3,
						exp_date =$4 where id = $5;`
	err := cp.DB.Raw(query, coupon.Code,
		coupon.DiscountPrice,
		coupon.MinimumPurchase,
		coupon.ExpDate, Id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (cp *CouponDatabase) DeleteCoupon(couponId uint) error {

	query := `delete from coupons where id = $1;`
	err := cp.DB.Exec(query, couponId).Error
	if err != nil {
		return err
	}
	return nil
}

func (cp *CouponDatabase) ViewCoupons() ([]domain.Coupon, error) {

	var body []domain.Coupon
	query := `select * from coupons;`
	err := cp.DB.Raw(query).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (cp *CouponDatabase) ViewCouponById(couponId uint) (domain.Coupon, error) {

	var body domain.Coupon
	query := `select * from coupons where id = $1`
	err := cp.DB.Raw(query, couponId).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (cp *CouponDatabase) ViewCouponByCode(code string) (domain.Coupon, error) {

	var body domain.Coupon
	query := `select * from coupons where code = $1;`
	err := cp.DB.Raw(query, code).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (c *CouponDatabase) ApplyCoupon(cid, uid uint) error {

	var body domain.UserCart
	query := `update user_carts set coupon_id = $1 where user_id = $2;`
	err := c.DB.Raw(query, cid, uid).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CouponDatabase) RemoveCoupon(uid uint) error {

	var body domain.UserCart
	query := `UPDATE user_carts     
			SET coupon_id = NULL  
			WHERE user_id = $1;`
	err := c.DB.Raw(query, uid).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CouponDatabase) FindCoupon(cid, uid uint) (domain.Coupon, error) {

	var body domain.Coupon
	query := `select * from user_carts where user_id = $1 and coupon_id= $2;`
	err := c.DB.Raw(query, uid, cid).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}
