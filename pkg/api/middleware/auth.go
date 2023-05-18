package middleware

import (
	"github.com/abhinandpn/project-ecom/pkg/helper"
	"github.com/gin-gonic/gin"
)

// Admin authentcation
func AuthUser(ctx *gin.Context) {
	helper.AuthHelperUser(ctx, "user")
}

// User authentication
func AuthAdmin(ctx *gin.Context) {
	helper.AuthHelperAdmin(ctx, "admin")
}
