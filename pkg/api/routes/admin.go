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
		login.POST("/", AdminHandler.AdminLogin)
		login.POST("/sudo", AdminHandler.SudoAdminLogin)
	}
	api.Use(middleware.AuthAdmin)
	{
		api.GET("/", AdminHandler.AdminHome) // Admin Home
		// user Side
		user := api.Group("/users")
		{
			user.GET("/", AdminHandler.Listuser)         // List all user
			user.PATCH("/block", AdminHandler.BlockUser) // Block User
		}
		// Product
		product := api.Group("/product")
		{
			product.GET("/all", ProductHandler.ListProducts) // list all product
			product.POST("/add", ProductHandler.AddProduct)  // Add Product
		}

	}
}
