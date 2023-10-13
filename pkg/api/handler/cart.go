package handler

import (
	"net/http"

	handlerInterface "github.com/abhinandpn/project-ecom/pkg/api/handler/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/helper"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"github.com/gin-gonic/gin"
)

type CartsHandler struct {
	CartUseCase service.CartUseCase
}

func NewCartHandler(cartUseCase service.CartUseCase) handlerInterface.CartHandler {

	return &CartsHandler{CartUseCase: cartUseCase}
}

func (c *CartsHandler) AddCart(ctx *gin.Context) {

	UID := helper.GetUserId(ctx)

	ParmId := ctx.Param("id")
	pfid, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find productinfoid",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	err = c.CartUseCase.AddToCart(UID, pfid, 1)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't add product in to cart",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "product Added",
		Data:       nil,
		Errors:     nil,
	})
}

func (c *CartsHandler) RemoveFromCart(ctx *gin.Context) {

	Uid := helper.GetUserId(ctx)

	ParmId := ctx.Param("id")
	pfid, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find productinfoid",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	err = c.CartUseCase.RemoveFromCart(Uid, pfid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't remove product from cart",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "product removed",
		Data:       nil,
		Errors:     nil,
	})
}

func (c *CartsHandler) ViewCart(ctx *gin.Context) {

	Uid := helper.GetUserId(ctx)

	// count, err1 := helper.StringToUInt(ctx.Query("count"))
	// pageNumber, err2 := helper.StringToUInt(ctx.Query("page_number"))

	// err1 = errors.Join(err1, err2)
	// if err1 != nil {
	// 	response := res.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
	// 	ctx.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// var Pagination req.ReqPagination
	// Pagination.Count = count
	// Pagination.PageNumber = pageNumber

	body, err := c.CartUseCase.CartDisplay(Uid)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't get cart",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "sucess to get all carts",
		Data:       body,
		Errors:     nil,
	})
}

func (c *CartsHandler) CartInfo(ctx *gin.Context) {

	Uid := helper.GetUserId(ctx)
	body, err := c.CartUseCase.CartInfo(Uid)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't get cart info",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "sucess to get all carts",
		Data:       body,
		Errors:     nil,
	})
}
