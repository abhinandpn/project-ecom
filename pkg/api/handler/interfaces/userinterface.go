package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {

	// user dashbord
	UserSignUp(ctx *gin.Context)         // signup user
	UserLogin(ctx *gin.Context)          // login user
	UserOtpLogin(ctx *gin.Context)       // otp logi for user with number
	UserLoginOtpVerify(ctx *gin.Context) // otp verify with auth
	UserHome(ctx *gin.Context)           // user home (verfication)
	UserInfo(ctx *gin.Context)           // user profile info(login details)
	UserLogout(ctx *gin.Context)         // user logout
	AddAddress(ctx *gin.Context)         // add address
	ListAllAddress(ctx *gin.Context)     // list all category - for user
	UpdateAddress(ctx *gin.Context)      // list all product - for user
	// CreateCart(ctx *gin.Context)
}
