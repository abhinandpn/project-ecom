package interfaces

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	ListProducts(ctx *gin.Context)
	AddProduct(ctx *gin.Context)
	EditProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	ViewProduct(ctx *gin.Context)
	Addcategory(ctx *gin.Context)
	EditCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
	Viewcategory(ctx *gin.Context)
	ViewFullcategory(ctx *gin.Context)
}
