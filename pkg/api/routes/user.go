package routes

import (
	"fmt"

	"github.com/abhinandpn/project-ecom/pkg/api/handler"
	"github.com/abhinandpn/project-ecom/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler, productHandler *handler.ProductHandler) {
	// login
	login := api.Group("/login")
	{
		fmt.Println("--------------- route------")
		login.POST("", userHandler.UserLogin)
		login.POST("/otp-send", userHandler.UserOtpLogin)
		login.POST("/otp-verify", userHandler.UserLoginOtpVerify)
	}

	// Signup
	signup := api.Group("/signup")
	{
		signup.POST("", userHandler.UserSignUp)
	}

	api.Use(middleware.AuthUser)
	{
		user := api.Group("/user")
		{
			user.GET("/home", userHandler.UserHome)      // User Home
			user.GET("/info", userHandler.UserHome)      // user Information
			user.POST("/logout", userHandler.UserLogout) // User Logout
		}
		// Address
		address := api.Group("address")
		{
			address.POST("/add", userHandler.AddAddress) // Add Address
		}
		// Product
		product := api.Group("/product")
		{
			product.GET("/all", productHandler.ListProducts) // List all product
		}

	}

}
