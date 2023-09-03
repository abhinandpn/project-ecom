package interfaces

import (
	"github.com/gin-gonic/gin"
)

type PaymentHandler interface {
	GetPaymentMethods(ctx *gin.Context)
	AddPaymentMethod(ctx *gin.Context)
	DeletePaymentMethod(ctx *gin.Context)

	// payment status
	CreatePaymentStatus(ctx *gin.Context)
	UpdatePaymentStatus(ctx *gin.Context)
	DeltePaymentStatus(ctx *gin.Context)
	FindPaymentStatusById(ctx *gin.Context)
	GetAllPaymentStatus(ctx *gin.Context)
}
