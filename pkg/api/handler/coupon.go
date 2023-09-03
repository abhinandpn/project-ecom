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

type CouponHandler struct {
	CouponUseCase service.CouponUseCase
}

func NewCouponHandler(couponUsecase service.CouponUseCase) handlerInterface.CouponHandler {

	return &CouponHandler{CouponUseCase: couponUsecase}
}

func (cp *CouponHandler) CrateCouponWithmoney(ctx *gin.Context) {

	var body req.CouponWithMoney
	if err := ctx.ShouldBindJSON(&body); err != nil {
		respones := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	err := cp.CouponUseCase.AddCoupon(body)
	if err != nil {
		response := res.ErrorResponse(400, "faild to add coupon", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	respones := res.SuccessResponse(200, "successfully create coupon", body)
	ctx.JSON(http.StatusOK, respones)
}

func (cp *CouponHandler) UpdateCoupon(ctx *gin.Context) {

	ParmId := ctx.Param("id")
	CouponId, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant get number",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var body req.CouponWithMoney
	err = ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	coupon, err := cp.CouponUseCase.UpdateCoupon(CouponId, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "updation failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully create coupon", coupon)
	ctx.JSON(http.StatusOK, respones)
}

func (cp *CouponHandler) DeleteCoupon(ctx *gin.Context) {

	ParmId := ctx.Param("id")
	CouponId, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant get number",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cp.CouponUseCase.DeleteCoupon(CouponId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant get number",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully delete coupon", CouponId)
	ctx.JSON(http.StatusOK, respones)
}

func (cp *CouponHandler) ViewCouponById(ctx *gin.Context) {

	ParmId := ctx.Param("id")
	CouponId, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant get id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	body, err := cp.CouponUseCase.ViewCouponById(CouponId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cat get coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully get coupon", body)
	ctx.JSON(http.StatusOK, respones)
}

func (cp *CouponHandler) ListCoupon(ctx *gin.Context) {

	body, err := cp.CouponUseCase.ListCoupon()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cat get coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully get coupon", body)
	ctx.JSON(http.StatusOK, respones)
}

func (c *CouponHandler) ViewCouponByCode(ctx *gin.Context) {

	Coupon := ctx.Param("name")

	body, err := c.CouponUseCase.ViewCouponByCode(Coupon)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cat get coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully get coupon", body)
	ctx.JSON(http.StatusOK, respones)
}

func (c *CouponHandler) ApplyCoupon(ctx *gin.Context) {

	userId := helper.GetUserId(ctx)
	Coupon := ctx.Param("name")
	err := c.CouponUseCase.ApplyCoupon(Coupon, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cat apply coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully apply coupon", Coupon)
	ctx.JSON(http.StatusOK, respones)
}

func (c *CouponHandler) RemoveCoupon(ctx *gin.Context) {

	userId := helper.GetUserId(ctx)
	Coupon := ctx.Param("name")
	err := c.CouponUseCase.RemoveCoupon(Coupon, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cat remove coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	respones := res.SuccessResponse(200, "successfully remove coupon", Coupon)
	ctx.JSON(http.StatusOK, respones)
}

func (c *CouponHandler) GetAppliedCoupon(ctx *gin.Context) {

	userId := helper.GetUserId(ctx)
	body, err := c.CouponUseCase.GetAppliedCoupon(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cat remove coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cat get coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	respones := res.SuccessResponse(200, "successfully get applied coupon", body)
	ctx.JSON(http.StatusOK, respones)
}
