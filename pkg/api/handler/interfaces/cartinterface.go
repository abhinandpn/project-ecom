package interfaces

import "github.com/gin-gonic/gin"

type CartHandler interface {

	// cart managment
	AddToCart(ctx *gin.Context) // product add to cart
	UserCart(ctx *gin.Context)  // user cart showing
}
