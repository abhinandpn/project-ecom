package routes

import (
	handlerInterface "github.com/abhinandpn/project-ecom/pkg/api/handler/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoute(api *gin.RouterGroup,
	AdminHandler handlerInterface.AdminHandler,
	ProductHandler handlerInterface.ProductHandler,
	orderHandler handlerInterface.OrderHandler,
	PaymentHandler handlerInterface.PaymentHandler,
	CouponHandler handlerInterface.CouponHandler,
) {
	// Sudo Admin Login Login
	login := api.Group("/login")
	{
		login.POST("", AdminHandler.AdminLogin)
		login.POST("/sudo", AdminHandler.SudoAdminLogin)
	}
	api.Use(middleware.AuthAdmin)
	{
		api.GET("", AdminHandler.AdminHome) // Admin Home
		// user Side
		user := api.Group("/users")
		{
			// user managment

			user.PATCH("/block/:id", AdminHandler.BlockUser) // Block User

			// Finding And Listing
			user.GET("", AdminHandler.Listuser)                              // List all user
			user.GET("/:id", AdminHandler.FindUserWithId)                    // find User With Id
			user.GET("/number/:number", AdminHandler.FindUserWithNumber)     // Find user with number
			user.GET("/email/:email", AdminHandler.FindUserWithEmail)        // Find User With email
			user.GET("/username/:username", AdminHandler.FindUserByUserName) // Find User With User name
			user.GET("/order/detail/:id", AdminHandler.UserOrderDetails)
			user.PATCH("/order/:id/:status", AdminHandler.ChangeOrderStatus)
		}
		// Product
		product := api.Group("/product")
		{
			product.GET("/all", ProductHandler.ListProducts)     // list all product
			product.GET("/:id", ProductHandler.ViewProduct)      // View Single product
			product.POST("/add", ProductHandler.AddProduct)      // Add Product
			product.PATCH("/:id", ProductHandler.UpdateProduct)  // Edit / update Product
			product.DELETE("/:id", ProductHandler.DeleteProduct) // Delete product

			// ----------- Sorting -----------
			product.GET("/all/colour", ProductHandler.ProductGetByColour)
			product.GET("/all/size", ProductHandler.ProductGetBySize)
			product.GET("/all/category", ProductHandler.ProductGetByCategory)
			product.GET("/all/brand", ProductHandler.ProductGetByBrand)
			product.GET("/all/name", ProductHandler.ProductGetByName)
			product.GET("/all/price", ProductHandler.ProductGetByPrice)
			product.GET("/all/quantity", ProductHandler.ProductGetByQuantity)
			product.GET("/get", ProductHandler.GetProductByString) // sort by string
		}
		// Category
		category := api.Group("/category")
		{
			category.GET("", ProductHandler.ViewFullcategory)  // View Categories
			category.GET("/:id", ProductHandler.Viewcategory)  // view Single Category
			category.POST("", ProductHandler.Addcategory)      // Add category
			category.PATCH("", ProductHandler.EditCategory)    // Edit category
			category.DELETE("", ProductHandler.DeleteCategory) // Delete category
		}
		// sub category
		subct := api.Group("/subcategory")
		{
			subct.GET("/all", ProductHandler.ViewFullSubCategory)  // list all sub category
			subct.POST("/add", ProductHandler.AddSubCategory)      // add sub category
			subct.DELETE("/:id", ProductHandler.DeleteSubCategory) // delete sub category
			subct.PATCH("/:id", ProductHandler.EditSubCategory)    // edit sub catagory
		}
		// brand
		brand := api.Group("/br/all/cand")
		{
			brand.POST("/add", ProductHandler.AddBrand)     // Add brand
			brand.DELETE("/:id", ProductHandler.DeletBrand) // Delete brand
			brand.GET("/all", ProductHandler.ViewBrands)    // View full brands
		}
		// payment method
		payment := api.Group("/payment/method")
		{
			payment.POST("/add/:name", PaymentHandler.AddPaymentMethod)       // add payment method
			payment.DELETE("/delete/:id", PaymentHandler.DeletePaymentMethod) // delete payment method
			payment.GET("/all", PaymentHandler.GetPaymentMethods)             // get all payment methods

		}
		// payment status
		status := api.Group("/payment/status")
		{
			status.POST("/add", PaymentHandler.CreatePaymentStatus)
			status.PATCH("/edit/:id", PaymentHandler.UpdatePaymentStatus)
			status.DELETE("/:id", PaymentHandler.DeltePaymentStatus)
			status.GET("/:id", PaymentHandler.FindPaymentStatusById)
			status.GET("/all", PaymentHandler.GetAllPaymentStatus)
		}
		// coupon
		coupon := api.Group("/coupon")
		{
			coupon.POST("/add/money", CouponHandler.CrateCouponWithmoney) // create coupon
			coupon.PATCH("/update/:id", CouponHandler.UpdateCoupon)       // update coupon
			coupon.DELETE("/:id", CouponHandler.DeleteCoupon)             // delete coupon
			coupon.GET("/code/:name", CouponHandler.ViewCouponByCode)     // get coupon by name
			coupon.GET("/:id", CouponHandler.ViewCouponById)              // get coupon by id
			coupon.GET("/all", CouponHandler.ListCoupon)                  // get full coupons

			//

		}
		// order
		orderstatus := api.Group("/order/status")
		{
			orderstatus.POST("/new", orderHandler.CreateOrderStatus)
			orderstatus.PATCH("/edit/:id", orderHandler.UpdateOrderStatus)
			orderstatus.DELETE("/:id", orderHandler.DeleteOrderStatus)
			orderstatus.GET("/all", orderHandler.GetAllOrderStatus)
			orderstatus.GET("/:id", orderHandler.FindOrderStatusById)
			orderstatus.GET("/get/:name", orderHandler.FindOrderStatusByStatus)

			// 01 - 09 - 2023 - Order status updation
			orderstatus.GET("/user/:id", orderHandler.ListAllOrderByUid)
			orderstatus.GET("/user/detail/:id", orderHandler.ListOrderDetailByUid)

			orderstatus.PATCH("/update/ordred/:id", orderHandler.OrderStatusToOrdered)
			orderstatus.PATCH("/update/delivered/:id", orderHandler.OrderStatusToDelivered)
			orderstatus.PATCH("/update/canclled/:id", orderHandler.OrderStatusToCancelled)
			orderstatus.PATCH("/update/returned/:id", orderHandler.OrderStatusToReturned)

		}

	}
}
