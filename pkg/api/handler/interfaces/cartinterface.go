package interfaces

import "github.com/gin-gonic/gin"

type CartHandler interface {
	AddToCart(ctx *gin.Context)
	UserCart(ctx *gin.Context)
}
