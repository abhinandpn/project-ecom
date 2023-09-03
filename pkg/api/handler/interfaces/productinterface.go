package interfaces

import "github.com/gin-gonic/gin"

type ProductHandler interface {

	// porduct
	AddProduct(ctx *gin.Context)    // add product
	UpdateProduct(ctx *gin.Context) // update product
	DeleteProduct(ctx *gin.Context) // delete product
	ListProducts(ctx *gin.Context)  // list all product
	ViewProduct(ctx *gin.Context)   // view each product

	// category
	Addcategory(ctx *gin.Context)      // add category
	EditCategory(ctx *gin.Context)     // edit category
	DeleteCategory(ctx *gin.Context)   // delete category
	Viewcategory(ctx *gin.Context)     // view each category
	ViewFullcategory(ctx *gin.Context) // list all category

	// brand
	AddBrand(ctx *gin.Context)
	DeletBrand(ctx *gin.Context)
	ViewBrands(ctx *gin.Context)

	// sub category
	AddSubCategory(ctx *gin.Context)
	DeleteSubCategory(ctx *gin.Context)
	ViewFullSubCategory(ctx *gin.Context)
	EditSubCategory(ctx *gin.Context)
	// ViewSubCategory(ctx *gin.Context)

	// ----------- Sorting -----------
	ProductGetByColour(ctx *gin.Context)
	ProductGetByCategory(ctx *gin.Context)
	ProductGetByName(ctx *gin.Context)
	ProductGetByBrand(ctx *gin.Context)
	ProductGetBySize(ctx *gin.Context)
	ProductGetByPrice(ctx *gin.Context)
	ProductGetByQuantity(ctx *gin.Context)
	GetProductByString(ctx *gin.Context)
}
