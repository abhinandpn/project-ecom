package routes

import (
	handlerInterface "github.com/abhinandpn/project-ecom/pkg/api/handler/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoute(api *gin.RouterGroup,
	AdminHandler handlerInterface.AdminHandler,
	ProductHandler handlerInterface.ProductHandler,
) {
	// Sudo Admin Login Login
	login := api.Group("/login")
	{
		login.POST("", AdminHandler.AdminLogin)
		login.POST("/sudo", AdminHandler.SudoAdminLogin)
	}
	api.Use(middleware.AuthAdmin)
	{
		api.GET("", AdminHandler.AdminHome) // Admin Home
		// user Side
		user := api.Group("/users")
		{
			// user managment

			user.PATCH("/block/:id", AdminHandler.BlockUser) // Block User

			// Finding And Listing
			user.GET("", AdminHandler.Listuser)                              // List all user
			user.GET("/:id", AdminHandler.FindUserWithId)                    // find User With Id
			user.GET("/number/:number", AdminHandler.FindUserWithNumber)     // Find user with number
			user.GET("/email/:email", AdminHandler.FindUserWithEmail)        // Find User With email
			user.GET("/username/:username", AdminHandler.FindUserByUserName) // Find User With User name
		}
		// Product
		product := api.Group("/product")
		{
			product.GET("", ProductHandler.ListProducts)         // list all product
			product.GET("/:id", ProductHandler.ViewProduct)      // View Single product
			product.POST("", ProductHandler.AddProduct)          // Add Product
			product.PATCH("/:id", ProductHandler.EditProduct)    // Edit / update Product
			product.DELETE("/:id", ProductHandler.DeleteProduct) // Delete product

			// ----------- Sorting -----------
			product.GET("/all/colour", ProductHandler.ProductGetByColour)
			product.GET("/all/size", ProductHandler.ProductGetBySize)
			product.GET("/all/category", ProductHandler.ProductGetByCategory)
			product.GET("/all/brand", ProductHandler.ProductGetByBrand)
			product.GET("/all/name", ProductHandler.ProductGetByName)
			product.GET("/all/price", ProductHandler.ProductGetByPrice)
			product.GET("/all/quantity", ProductHandler.ProductGetByQuantity)
		}
		// Category
		category := api.Group("/category")
		{
			category.GET("", ProductHandler.ViewFullcategory)  // View Categories
			category.GET("/:id", ProductHandler.Viewcategory)  // view Single Category
			category.POST("", ProductHandler.Addcategory)      // Add category
			category.PATCH("", ProductHandler.EditCategory)    // Edit category
			category.DELETE("", ProductHandler.DeleteCategory) // Delete category
		}
		// sub category
		subct := api.Group("/subcategory")
		{
			subct.GET("/all", ProductHandler.ViewFullSubCategory)  // list all sub category
			subct.POST("/add", ProductHandler.AddSubCategory)      // add sub category
			subct.DELETE("/:id", ProductHandler.DeleteSubCategory) // delete sub category
			subct.PATCH("/:id", ProductHandler.EditSubCategory)    // edit sub catagory
		}
		// brand
		brand := api.Group("/brand")
		{
			brand.POST("/add", ProductHandler.AddBrand)     // Add brand
			brand.DELETE("/:id", ProductHandler.DeletBrand) // Delete brand
			brand.GET("/all", ProductHandler.ViewBrands)    // View full brands
		}
	}
}
