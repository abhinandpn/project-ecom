package handler

import (
	"net/http"

	"github.com/abhinandpn/project-ecom/pkg/helper"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"github.com/gin-gonic/gin"
)

type CartsHandler struct {
	CartUseCase service.CartUseCase
}

func NewCartHandler(cartUseCase service.CartUseCase) *CartsHandler {

	return &CartsHandler{CartUseCase: cartUseCase}
}

// add to cart
func (cr *CartsHandler) AddToCart(ctx *gin.Context) {

	UserId := helper.GetUserId(ctx)

	ParmId := ctx.Param("id")

	Pid, err := helper.StringToUInt(ParmId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	err = cr.CartUseCase.AddToCart(ctx, Pid, UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Add to cart failed",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "product Added to cart",
		Data:       nil,
		Errors:     nil,
	})
}

// list Usercart
func (cr *CartsHandler) UserCart(ctx *gin.Context) {

	UserId := helper.GetUserId(ctx)

	var UserCart res.CartRes
	// if err := ctx.ShouldBindJSON(&UserCart); err != nil {
	// 	response := res.ErrorResponse(400, "invalid input", err.Error(), UserCart)
	// 	ctx.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	UserCart, err := cr.CartUseCase.UserCart(ctx, UserId)
	if err != nil {
		response := res.ErrorResponse(400, "failed to find user cart", err.Error(), UserCart)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := res.SuccessResponse(200, "successfully product updated ", UserCart)
	ctx.JSON(200, response)
}

// remove from cart
// List cart items
