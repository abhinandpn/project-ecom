package interfaces

import "github.com/gin-gonic/gin"

type ProductHandler interface {

	// porduct
	AddProduct(ctx *gin.Context)    // add product
	EditProduct(ctx *gin.Context)   // edit product
	DeleteProduct(ctx *gin.Context) // delete product
	ListProducts(ctx *gin.Context)  // list all product
	ViewProduct(ctx *gin.Context)   // view each product

	// category
	Addcategory(ctx *gin.Context)      // add category
	EditCategory(ctx *gin.Context)     // edit category
	DeleteCategory(ctx *gin.Context)   // delete category
	Viewcategory(ctx *gin.Context)     // view each category
	ViewFullcategory(ctx *gin.Context) // list all category
}