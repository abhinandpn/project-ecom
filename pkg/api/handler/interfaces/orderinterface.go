package interfaces

import "github.com/gin-gonic/gin"

type OrderHandler interface {
	BuyNow(ctx *gin.Context)
}
