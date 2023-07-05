package interfaces

import (
	"github.com/gin-gonic/gin"
)

type PaymentHandler interface {
	GetPaymentMethods(ctx *gin.Context)
	AddPaymentMethod(ctx *gin.Context)
	DeletePaymentMethod(ctx *gin.Context)
}
