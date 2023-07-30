package interfaces

import "github.com/gin-gonic/gin"

type CouponHandler interface {
	CrateCoupon(ctx *gin.Context)
	UpdateCoupon(ctx *gin.Context)
	DeleteCoupon(ctx *gin.Context)
	ViewCouponById(ctx *gin.Context)
	ListCoupon(ctx *gin.Context)
	ViewCouponByCode(ctx *gin.Context)
}
