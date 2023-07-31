package interfaces

import "github.com/gin-gonic/gin"

type CouponHandler interface {
	CrateCouponWithmoney(ctx *gin.Context)
	UpdateCoupon(ctx *gin.Context)
	DeleteCoupon(ctx *gin.Context)
	ViewCouponById(ctx *gin.Context)
	ListCoupon(ctx *gin.Context)
	ViewCouponByCode(ctx *gin.Context)

	//
	ApplyCoupon(ctx *gin.Context)
	RemoveCoupon(ctx *gin.Context)
}
