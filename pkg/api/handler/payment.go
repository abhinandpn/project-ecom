package handler

import (
	"net/http"

	handlerInterface "github.com/abhinandpn/project-ecom/pkg/api/handler/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/helper"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	PaymentUseCase service.PaymentuseCase
}

func NewPaymentHandler(paymentUseCase service.PaymentuseCase) handlerInterface.PaymentHandler {

	return &PaymentHandler{PaymentUseCase: paymentUseCase}
}

func (p *PaymentHandler) GetPaymentMethods(ctx *gin.Context) {

	body, err := p.PaymentUseCase.PaymentMethods()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find payment methods",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	if body == nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "payment method is empty",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully get payment methods", body)
	ctx.JSON(http.StatusOK, respones)
}

func (p *PaymentHandler) AddPaymentMethod(ctx *gin.Context) {

	name := ctx.Param("name")
	status := true
	err := p.PaymentUseCase.AddPaymentMethod(name, status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't add payment methods",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully add payment methods", name)
	ctx.JSON(http.StatusOK, respones)
}

func (p *PaymentHandler) DeletePaymentMethod(ctx *gin.Context) {

	ParamId := ctx.Param("id")
	id, err := helper.StringToUInt(ParamId)
	if err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	err = p.PaymentUseCase.DeletePaymentMethod(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't delete payment methods",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully delete payment methods", id)
	ctx.JSON(http.StatusOK, respones)
}

func (p *PaymentHandler) CreatePaymentStatus(ctx *gin.Context) {

	var body req.NewPaymentStatus
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		respones := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	err = p.PaymentUseCase.CreatePaymentStatus(body.Status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't create payment status",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully create payment status", body)
	ctx.JSON(http.StatusOK, respones)
}

func (p *PaymentHandler) UpdatePaymentStatus(ctx *gin.Context) {

	ParamId := ctx.Param("id")
	id, err := helper.StringToUInt(ParamId)
	if err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	var body req.NewPaymentStatus
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		respones := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}
	status, err := p.PaymentUseCase.UpdatePaymentStatus(id, body.Status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't update payment status",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully update payment status", status)
	ctx.JSON(http.StatusOK, respones)
}

func (p *PaymentHandler) DeltePaymentStatus(ctx *gin.Context) {

	ParamId := ctx.Param("id")
	id, err := helper.StringToUInt(ParamId)
	if err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	err = p.PaymentUseCase.DeltePaymentStatus(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't delete payment status",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully update payment status", id)
	ctx.JSON(http.StatusOK, respones)
}

func (p *PaymentHandler) FindPaymentStatusById(ctx *gin.Context) {

	ParamId := ctx.Param("id")
	id, err := helper.StringToUInt(ParamId)
	if err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	body, err := p.PaymentUseCase.FindPaymentStatusById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't delete payment status",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully get payment status", body)
	ctx.JSON(http.StatusOK, respones)
}

func (p *PaymentHandler) GetAllPaymentStatus(ctx *gin.Context) {
	body, err := p.PaymentUseCase.GetAllPaymentStatus()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't get payment status",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully get payment status", body)
	ctx.JSON(http.StatusOK, respones)
}
