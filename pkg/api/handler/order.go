package handler

import (
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

	ParmId := ctx.Param("id")
	payid, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	err = o.orderUseCase.CartOrderAll(Uid, payid)
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
