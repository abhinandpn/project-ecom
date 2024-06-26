package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {

	// user dashbord
	UserHome(ctx *gin.Context)           // user home (verfication)
	UserInfo(ctx *gin.Context)           // user profile info(login details)
	UserLogout(ctx *gin.Context)         // user logout
	UserLogin(ctx *gin.Context)          // login user
	UserSignUp(ctx *gin.Context)         // signup user
	UserOtpLogin(ctx *gin.Context)       // otp logi for user with number
	UserLoginOtpVerify(ctx *gin.Context) // otp verify with auth
	AddAddress(ctx *gin.Context)         // add address
	UserStatus(ctx *gin.Context)         // Get user Block status
	// address
	ListAllAddress(ctx *gin.Context)        // List all category - for user
	UpdateAddress(ctx *gin.Context)         // List all product - for user
	GetUserDefaultAddress(ctx *gin.Context) // Get user default address
	MakeAddressDefault(ctx *gin.Context)    // Make address Default

	// wishlist
	AddIntoWishlit(ctx *gin.Context)     // add product in to wishlist
	RemoveFromWIshList(ctx *gin.Context) // remove product from wishlist
	ViewWishList(ctx *gin.Context)       // view wishlist
}
