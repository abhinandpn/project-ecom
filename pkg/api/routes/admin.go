package routes

import (
	"github.com/abhinandpn/project-ecom/pkg/api/handler"
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
}
