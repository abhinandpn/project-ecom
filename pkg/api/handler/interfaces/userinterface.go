package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {
	UserSignUp(ctx *gin.Context)
	UserLogin(ctx *gin.Context)
	UserOtpLogin(ctx *gin.Context)
	UserLoginOtpVerify(ctx *gin.Context)
	UserHome(ctx *gin.Context)
	UserInfo(ctx *gin.Context)
	UserLogout(ctx *gin.Context)
	AddAddress(ctx *gin.Context)
	ListAllAddress(ctx *gin.Context)
	UpdateAddress(ctx *gin.Context)
	// CreateCart(ctx *gin.Context)
}
