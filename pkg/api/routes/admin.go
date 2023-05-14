package routes

import (
	"github.com/abhinandpn/project-ecom/pkg/api/handler"
	"github.com/abhinandpn/project-ecom/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoute(api *gin.RouterGroup,
	AdminHandler *handler.AdminHandler,
) {
	// Login
	login := api.Group("/admin")
	{
		login.POST("/", AdminHandler.AdminLogin)
	}
	api.Use(middleware.AuthAdmin)
	{
		api.GET("/", AdminHandler.AdminHome)
		// user Side
		user := api.Group("/users")
		user.GET("/", AdminHandler.Listuser)
		user.PATCH("/block", AdminHandler.BlockUser)
	}
}
