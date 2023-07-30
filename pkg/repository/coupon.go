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

func (cp *CouponDatabase) AddCoupon(coupon req.CouponReq) error {

	var body domain.Coupon
	query := `insert into coupons (code,
					discount_persentage,
					discount_price,
					minimum_purchase,
					exp_date)values ($1,$2,$3,$4,$5);`
	err := cp.DB.Raw(query,
		coupon.Code,
		coupon.DiscountPersentage,
		coupon.DiscountPrice,
		coupon.MinimumPurchase).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (cp *CouponDatabase) UpdateCoupon(coupon req.CouponReq, Id uint) (domain.Coupon, error) {

	var body domain.Coupon
	query := `update coupons set code = $1,
						discount_persentage =$2 ,
						discount_price = $3,
						minimum_purchase = $4,
						exp_date =$5 where id = $7;`
	err := cp.DB.Raw(query, coupon.Code,
		coupon.DiscountPersentage,
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
