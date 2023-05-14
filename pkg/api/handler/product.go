package handler

import (
	"errors"
	"net/http"

	"github.com/abhinandpn/project-ecom/pkg/helper"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductuseCase service.ProductuseCase
}

func NewProductHandler(productUsecase service.ProductuseCase) *ProductHandler {
	return &ProductHandler{ProductuseCase: productUsecase}
}
func (pr *ProductHandler) ListProducts(ctx *gin.Context) {
	count, err1 := helper.StringToUInt(ctx.Query("count"))
	pageNumber, err2 := helper.StringToUInt(ctx.Query("page_number"))

	err1 = errors.Join(err1, err2)
	if err1 != nil {
		response := res.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	pagination := req.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}

	products, err := pr.ProductuseCase.GetProducts(ctx, req.PageNation(pagination))

	if err != nil {
		response := res.ErrorResponse(500, "faild to get all products", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if products == nil {
		response := res.SuccessResponse(200, "there is no products to show", nil)
		ctx.JSON(http.StatusOK, response)
		return
	}

	respones := res.SuccessResponse(200, "successfully got all products", products)
	ctx.JSON(http.StatusOK, respones)
}
