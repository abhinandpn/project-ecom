package handler

import (
	"errors"
	"fmt"
	"net/http"

	handlerInterface "github.com/abhinandpn/project-ecom/pkg/api/handler/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/helper"
	services "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase services.OrderUseCase
}

func NewOrderHandler(usecase services.OrderUseCase) handlerInterface.OrderHandler {
	return &OrderHandler{
		orderUseCase: usecase,
	}
}
func (o *OrderHandler) BuyNow(ctx *gin.Context) {

	Uid := helper.GetUserId(ctx)

	ParmId := ctx.Param("id")
	Pfid, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	err = o.orderUseCase.OrderByPfId(Uid, Pfid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "buy now failed",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	respones := res.SuccessResponse(200, "successfully product ordered", nil)
	ctx.JSON(http.StatusOK, respones)
}

func (o *OrderHandler) CartAllOrder(ctx *gin.Context) {

	Uid := helper.GetUserId(ctx)
	PaymentId, err1 := helper.StringToUInt(ctx.Query("payid"))
	AddressId, err2 := helper.StringToUInt(ctx.Query("addressid"))

	var copid uint

	err := errors.Join(err1, err2)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	err = o.orderUseCase.CartOrderAll(Uid, PaymentId, copid, AddressId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't order cart products",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully cart product ordered", nil)
	ctx.JSON(http.StatusOK, respones)
}

func (o *OrderHandler) CartOrderStatus(ctx *gin.Context) {

	Uid := helper.GetUserId(ctx)

	ParmId := ctx.Param("id")
	Oid, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	body, err := o.orderUseCase.OrderStatus(Uid, Oid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find order status",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully order status", body)
	ctx.JSON(http.StatusOK, respones)
}

func (o *OrderHandler) OrderByproductId(ctx *gin.Context) {
	// CreateOrderStatus(ctx * gin.Context)
	Uid := helper.GetUserId(ctx)

	var body req.OrderByProduct
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// updation for body
	body.UserId = Uid
	// body.AddressId
	// body.CouponId
	// body.PaymentMethodId
	// body.ProductInfoId
	// body.Quantity
}

func (o *OrderHandler) OrderDetail(ctx *gin.Context) {

	Uid := helper.GetUserId(ctx)

	body, err := o.orderUseCase.UserOrders(Uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't get order details",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully get order details", body)
	ctx.JSON(http.StatusOK, respones)

}

func (or *OrderHandler) CreateOrderStatus(ctx *gin.Context) {

	var body req.OrderStatus
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		respones := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	err = or.orderUseCase.CreateOrderStatus(body.Status)
	if err != nil {
		respones := res.ErrorResponse(400, "cant create order status", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	respones := res.SuccessResponse(200, "successfully create order status", body)
	ctx.JSON(http.StatusOK, respones)
}

func (or *OrderHandler) UpdateOrderStatus(ctx *gin.Context) {

	ParmId := ctx.Param("id")
	Oid, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find order status id",
			Errors:     err.Error(),
			Data:       nil,
		})
		return

	}
	var body req.OrderStatus
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		respones := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	order, err := or.orderUseCase.UpdateOrderStatus(Oid, body.Status)
	if err != nil {
		respones := res.ErrorResponse(400, "cant update order status", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	respones := res.SuccessResponse(200, "successfully update order status", order)
	ctx.JSON(http.StatusOK, respones)
}

func (or *OrderHandler) DeleteOrderStatus(ctx *gin.Context) {

	ParmId := ctx.Param("id")
	Oid, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find order status id",
			Errors:     err.Error(),
			Data:       nil,
		})
		return

	}
	err = or.orderUseCase.DeletOrderStatus(Oid)
	if err != nil {
		respones := res.ErrorResponse(400, "cant delete order status", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	respones := res.SuccessResponse(200, "successfully delet order status", nil)
	ctx.JSON(http.StatusOK, respones)
}

func (or *OrderHandler) FindOrderStatusByStatus(ctx *gin.Context) {

	Status := ctx.Param("name")
	body, err := or.orderUseCase.FindOrderStatusByStatus(Status)
	if err != nil {
		respones := res.ErrorResponse(400, "cant find order status", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	respones := res.SuccessResponse(200, "successfully get order status", body)
	ctx.JSON(http.StatusOK, respones)
}

func (or *OrderHandler) FindOrderStatusById(ctx *gin.Context) {

	ParmId := ctx.Param("id")
	Oid, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find order status id",
			Errors:     err.Error(),
			Data:       nil,
		})
		return

	}
	body, err := or.orderUseCase.FindOrderStatusById(Oid)
	if err != nil {
		respones := res.ErrorResponse(400, "cant find order status", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	respones := res.SuccessResponse(200, "successfully get order status", body)
	ctx.JSON(http.StatusOK, respones)
}

func (or *OrderHandler) GetAllOrderStatus(ctx *gin.Context) {

	body, err := or.orderUseCase.FindAllOrderStatus()
	if err != nil {
		respones := res.ErrorResponse(400, "cant find order status", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	respones := res.SuccessResponse(200, "successfully get order status", body)
	ctx.JSON(http.StatusOK, respones)
}

func (o *OrderHandler) UpdatedCartAllOrder(ctx *gin.Context) {

	Uid := helper.GetUserId(ctx)
	PaymentId, err1 := helper.StringToUInt(ctx.Query("payid"))
	AddressId, err2 := helper.StringToUInt(ctx.Query("addressid"))

	err := errors.Join(err1, err2)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	err = o.orderUseCase.UpdatedCartAllOrder(Uid, PaymentId, AddressId)
	if err != nil {
		respones := res.ErrorResponse(400, "cant order cart product", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	respones := res.SuccessResponse(200, "successfully  order status", Uid)
	ctx.JSON(http.StatusOK, respones)
}

// 01 - 09 - 2023 - Order status updation
func (o *OrderHandler) ListAllOrderByUid(ctx *gin.Context) {

	ParmId := ctx.Param("id")
	Uid, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find order status id",
			Errors:     err.Error(),
			Data:       nil,
		})
		return

	}
	fmt.Println("==== > > > > > ", Uid)
	orders, err := o.orderUseCase.ListAllOrderByUid(Uid)
	if err != nil {
		respones := res.ErrorResponse(400, "can't get orders", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	respones := res.SuccessResponse(200, "successfully get orders", orders)
	ctx.JSON(http.StatusOK, respones)
}

func (o *OrderHandler) OrderStatusToOrdered(ctx *gin.Context) {

	ParmId := ctx.Param("id")
	uid, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find order status id",
			Errors:     err.Error(),
			Data:       nil,
		})
		return

	}
	err = o.orderUseCase.OrderStatusToOrdered(uid)
	if err != nil {
		respones := res.ErrorResponse(400, "can't update status to ordered", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	respones := res.SuccessResponse(200, "successfully update status to orderd", nil)
	ctx.JSON(http.StatusOK, respones)
}

func (o *OrderHandler) OrderStatusToDelivered(ctx *gin.Context) {

	ParmId := ctx.Param("id")
	uid, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find order status id",
			Errors:     err.Error(),
			Data:       nil,
		})
		return

	}
	err = o.orderUseCase.OrderStatusToDelivered(uid)
	if err != nil {
		respones := res.ErrorResponse(400, "can't update status to delivered", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	respones := res.SuccessResponse(200, "successfully update status to delivred", nil)
	ctx.JSON(http.StatusOK, respones)
}

func (o *OrderHandler) OrderStatusToCancelled(ctx *gin.Context) {

	ParmId := ctx.Param("id")
	uid, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find order status id",
			Errors:     err.Error(),
			Data:       nil,
		})
		return

	}
	err = o.orderUseCase.OrderStatusToCancelled(uid)
	if err != nil {
		respones := res.ErrorResponse(400, "can't update status to calclled", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	respones := res.SuccessResponse(200, "successfully update status to calclled", nil)
	ctx.JSON(http.StatusOK, respones)
}

func (o *OrderHandler) OrderStatusToReturned(ctx *gin.Context) {

	ParmId := ctx.Param("id")
	uid, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find order status id",
			Errors:     err.Error(),
			Data:       nil,
		})
		return

	}
	err = o.orderUseCase.OrderStatusToReturned(uid)
	if err != nil {
		respones := res.ErrorResponse(400, "can't update status to returned", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	respones := res.SuccessResponse(200, "successfully update status to returned", nil)
	ctx.JSON(http.StatusOK, respones)
}

func (o *OrderHandler) ListOrderDetailByUid(ctx *gin.Context) {

	ParmId := ctx.Param("id")
	uid, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find order status id",
			Errors:     err.Error(),
			Data:       nil,
		})
		return

	}
	fmt.Println("user id( handler ) --- > ", uid)
	body, err := o.orderUseCase.ListOrderDetailByUid(uid)
	if err != nil {
		respones := res.ErrorResponse(400, "can't get order details", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	respones := res.SuccessResponse(200, "successfully update status to returned", body)
	ctx.JSON(http.StatusOK, respones)
}
