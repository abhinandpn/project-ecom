package middleware

import (
	"net/http"
	"time"

	"github.com/abhinandpn/project-ecom/pkg/auth"
	"github.com/gin-gonic/gin"
)

// Admin authentcation
func AuthUser(ctx *gin.Context) {
	authHelper(ctx, "user")
}

// User authentication
func AuthAdmin(ctx *gin.Context) {
	authHelper(ctx, "admin")
}

// helper to get cookie and validate the token and expire time
func authHelper(ctx *gin.Context, user string) {

	tokenString, err := ctx.Cookie(user + "-auth") // get cookie for user or admin with name
	if err != nil || tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized User Please Login",
		})
		return
	}

	claims, err := auth.ValidateToken(tokenString) // auth function validate the token and return claims with error
	if err != nil || tokenString == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "Unauthorized User Please Login",
		})
		return
	}

	// check the cliams expire time
	if time.Now().Unix() > claims.ExpiresAt {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": 401,
			"msg":        "User Need Re-Login time expired",
		})
		return
	}

	// claim the userId and set it on context
	ctx.Set("userId", claims.Id)
}
