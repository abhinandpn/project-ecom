package interfaces

import "github.com/gin-gonic/gin"

type OrderHandler interface {
	BuyNow(ctx *gin.Context)
	CartAllOrder(ctx *gin.Context)
	CartOrderStatus(ctx *gin.Context)
	OrderByproductId(ctx *gin.Context)
	OrderDetail(ctx *gin.Context)

	// ----------------------
	CreateOrderStatus(ctx *gin.Context)
	UpdateOrderStatus(ctx *gin.Context)
	DeleteOrderStatus(ctx *gin.Context)
	FindOrderStatusByStatus(ctx *gin.Context)
	FindOrderStatusById(ctx *gin.Context)
	GetAllOrderStatus(ctx *gin.Context)
	// ----------------------
	UpdatedCartAllOrder(ctx *gin.Context)

	// 01 - 09 - 2023 - Order status updation
	ListOrderDetailByUid(ctx *gin.Context)
	ListAllOrderByUid(ctx *gin.Context)
	OrderStatusToOrdered(ctx *gin.Context)
	OrderStatusToDelivered(ctx *gin.Context)
	OrderStatusToCancelled(ctx *gin.Context)
	OrderStatusToReturned(ctx *gin.Context)
}
