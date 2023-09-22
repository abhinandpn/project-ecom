package routes

import (
	handlerInterface "github.com/abhinandpn/project-ecom/pkg/api/handler/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup,
	userHandler handlerInterface.UserHandler,
	productHandler handlerInterface.ProductHandler,
	cartHandler handlerInterface.CartHandler,
	orderHandler handlerInterface.OrderHandler,
	paymentHandler handlerInterface.PaymentHandler,
	CouponHandler handlerInterface.CouponHandler) {

	// login
	login := api.Group("/login")
	{
		login.POST("", userHandler.UserLogin)
		login.POST("/otp-send", userHandler.UserOtpLogin)
		login.POST("/otp-verify", userHandler.UserLoginOtpVerify)
	}

	// Signup
	signup := api.Group("/signup")
	{
		signup.POST("", userHandler.UserSignUp)
	}
	product := api.Group("/product")
	{
		// ----------- Sorting -----------
		product.GET("/all/colour", productHandler.ProductGetByColour)
		product.GET("/all/size", productHandler.ProductGetBySize)
		product.GET("/all/category", productHandler.ProductGetByCategory)
		product.GET("/all/brand", productHandler.ProductGetByBrand)
		product.GET("/all/name", productHandler.ProductGetByName)
		product.GET("/all/price", productHandler.ProductGetByPrice)
		product.GET("/get/:name", productHandler.GetProductByString)

		// List all product
		product.GET("/all", productHandler.ListProducts)
	}
	api.Use(middleware.AuthUser)
	{
		// logout
		logout := api.Group("/logout")
		{
			logout.DELETE("", userHandler.UserLogout)
		}

		api.Use(userHandler.UserStatus)
		{

			user := api.Group("/user")
			{
				user.GET("/home", userHandler.UserHome)      // User Home
				user.GET("/info", userHandler.UserInfo)      // user Information
				user.POST("/logout", userHandler.UserLogout) // User Logout
			}
			// Address
			address := api.Group("/address")
			{
				address.POST("/add", userHandler.AddAddress)                  // Add Address
				address.GET("/all", userHandler.ListAllAddress)               // List all Address
				address.PATCH("/update", userHandler.UpdateAddress)           // Update Address
				address.GET("/default", userHandler.GetUserDefaultAddress)    // get default address
				address.PATCH("/default/:id", userHandler.MakeAddressDefault) // make address default
			}

			// Category
			category := api.Group("/category")
			{
				category.GET("", productHandler.ViewFullcategory) // View full category
			}
			// Cart
			cart := api.Group("/cart")
			{
				cart.POST("/:id", cartHandler.AddCart)          // product ad to cart
				cart.DELETE("/:id", cartHandler.RemoveFromCart) // product remove from cart
				cart.GET("/all", cartHandler.ViewCart)          // view all cart
				cart.GET("/info", cartHandler.CartInfo)         // cart info
			}
			//order
			order := api.Group("/order")
			{
				order.POST("/cart", orderHandler.CartAllOrder)            // cart all product order
				order.POST("/new/cart", orderHandler.UpdatedCartAllOrder) // updated
				order.GET("/status/:id", orderHandler.CartOrderStatus)    // order status
				order.POST("/buynow/:id", orderHandler.BuyNow)            // order by pfid
				order.GET("/detail", orderHandler.OrderDetail)            // order details
			}
			// wishlist
			wishlist := api.Group("/wishlist")
			{
				wishlist.POST("/add/:id", userHandler.AddIntoWishlit)          // add product in to wishlist
				wishlist.DELETE("/remove/:id", userHandler.RemoveFromWIshList) // remove product in to wishlist
				wishlist.GET("/all", userHandler.ViewWishList)                 // view wishlist
			}
			// payment
			payment := api.Group("payment")
			{
				payment.GET("/methods", paymentHandler.GetPaymentMethods) // get payment mehods
			}
			// coupon
			coupon := api.Group("/coupon")
			{
				coupon.PATCH("/apply/:name", CouponHandler.ApplyCoupon)   // apply coupon
				coupon.PATCH("/remove/:name", CouponHandler.RemoveCoupon) // remove coupon
				coupon.GET("/show", CouponHandler.GetAppliedCoupon)       // get applied coupon
			}
		}
	}
}
