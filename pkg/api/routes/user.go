package routes

import (
	"github.com/abhinandpn/project-ecom/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler) {
	// login
	login := api.Group("/login")
	{
		login.POST("/", userHandler.UserLogin)
	}
	// Signup
	signup := api.Group("/signup")
	{
		signup.POST("/", userHandler.UserSignUp)
	}
}
