package routes

import (
	"github.com/abhinandpn/project-ecom/pkg/api/handler"
	"github.com/abhinandpn/project-ecom/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoute(api *gin.RouterGroup,
	AdminHandler *handler.AdminHandler,
	ProductHandler *handler.ProductHandler,
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
			user.PATCH("/block", AdminHandler.BlockUser) // Block User

			// Finding And Listing
			user.GET("", AdminHandler.Listuser)                     // List all user
			user.GET("/:id", AdminHandler.FindUserWithId)           // find User With Id
			user.GET("/:number", AdminHandler.FindUserWithNumber)   // Find user with number
			user.GET("/:email", AdminHandler.FindUserWithEmail)     // Find User With email
			user.GET("/:username", AdminHandler.FindUserByUserName) // Find User With User name
		}
		// Product
		product := api.Group("/product")
		{
			product.GET("", ProductHandler.ListProducts)     // list all product
			product.GET("/:id", ProductHandler.ViewProduct)  // View Single product
			product.POST("", ProductHandler.AddProduct)      // Add Product
			product.PATCH("", ProductHandler.EditProduct)    // Edit / update Product
			product.DELETE("", ProductHandler.DeleteProduct) // Delete product
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

	}
}
