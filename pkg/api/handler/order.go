package handler

import (
	"net/http"

	handlerInterface "github.com/abhinandpn/project-ecom/pkg/api/handler/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/helper"
	services "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
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
