package interfaces

import "github.com/gin-gonic/gin"

type CartHandler interface {

	// Add cart
	AddCart(ctx *gin.Context)

	// Remove Cart
	RemoveFromCart(ctx *gin.Context)

	// List cart
	ViewCart(ctx *gin.Context)
}
