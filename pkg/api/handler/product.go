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

//
// Product

// ListProducts-Admin godoc
// @summary api for admin to show products
// @security ApiKeyAuth
// @tags Admin Products
// @id ListProducts-Admin
// @Param page_number query int false "Page Number"
// @Param count query int false "Count Of Order"
// @Router /admin/products [get]
// @Success 200 {object} res.Response{} "successfully got all products"
// @Failure 500 {object} res.Response{}  "faild to get all products"

// ListProducts-User godoc
// @summary api for user to show products
// @security ApiKeyAuth
// @tags User Products
// @id ListProducts-User
// @Param page_number query int false "Page Number"
// @Param count query int false "Count Of Order"
// @Router /products [get]
// @Success 200 {object} res.Response{} "successfully got all products"
// @Failure 500 {object} res.Response{}  "faild to get all products"
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

	products, err := pr.ProductuseCase.ViewFullProduct(ctx, req.PageNation(pagination))

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

// AddProducts godoc
// @summary api for admin to update a product
// @id AddProducts
// @tags Admin Products
// @Param input body req.ReqProduct{} true "inputs"
// @Router /admin/products [post]
// @Success 200 {object} res.Response{} "successfully product added"
// @Failure 400 {object} res.Response{} "invalid input"
func (pr *ProductHandler) AddProduct(ctx *gin.Context) {

	var body req.ReqProduct

	if err := ctx.ShouldBindJSON(&body); err != nil {
		respones := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	// product := domain.Product{
	// 	ProductName: body.ProductName,
	// 	CategoryID:  body.CategoryID,
	// 	Discription: body.Discription,
	// 	Price:       body.Price,
	// 	Info: domain.ProductInfo{
	// 		Colour: body.Color,
	// 		Brand:  body.Brand,
	// 		Size:   body.Size,
	// 	},
	// }

	err := pr.ProductuseCase.AddProduct(ctx, body)

	if err != nil {
		response := res.ErrorResponse(400, "faild to add product", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := res.SuccessResponse(200, "successfully product added", body)
	ctx.JSON(http.StatusOK, response)

}

func (pr *ProductHandler) EditProduct(ctx *gin.Context) {

	var body req.ReqProduct
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ParamId := ctx.Param("id")
	id, err := helper.StringToUInt(ParamId)
	if err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), id)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = pr.ProductuseCase.UpdateProduct(ctx, body, id)
	if err != nil {
		response := res.ErrorResponse(400, "faild to update product", err.Error(), body)
		ctx.JSON(400, response)
		return
	}

	response := res.SuccessResponse(200, "successfully product updated ", body)
	ctx.JSON(200, response)
}

func (pr *ProductHandler) DeleteProduct(ctx *gin.Context) {

	ParmId := ctx.Param("id")

	id, err := helper.StringToUInt(ParmId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}

	err = pr.ProductuseCase.DeleteProduct(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't delete product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "product deleted",
		Data:       nil,
		Errors:     nil,
	})
}

func (pr *ProductHandler) ViewProduct(ctx *gin.Context) {

	ParmId := ctx.Param("id")
	id, err := helper.StringToUInt(ParmId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	data, err := pr.ProductuseCase.FindProductById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't Find product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "product Found",
		Data:       data,
		Errors:     nil,
	})
}

// Category Handler

// ----------------AddCategory----------------

func (ct *ProductHandler) Addcategory(ctx *gin.Context) {

	var body res.CategoryRes

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	name := body.CategoryName

	category, err := ct.ProductuseCase.AddCategory(ctx, name)

	if err != nil {
		response := res.ErrorResponse(400, "failed to add category", err.Error(), category)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := res.SuccessResponse(200, "successfully category added", body)
	ctx.JSON(200, response)

}

// ----------------EditCategory----------------
func (ct *ProductHandler) EditCategory(ctx *gin.Context) {

	// check the category and find the json binding
	var body req.UpdateCategoryReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// update category
	err := ct.ProductuseCase.UpdateCategory(ctx, body)
	if err != nil {
		response := res.ErrorResponse(400, "unable to update category", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	// response
	response := res.SuccessResponse(200, "successfully category updated", body)
	ctx.JSON(200, response)
}

// ----------------DeleteCategory----------------

func (ct *ProductHandler) DeleteCategory(ctx *gin.Context) {

	var body req.CategoryReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := ct.ProductuseCase.DeleteCategory(ctx, body.CategoryName)
	if err != nil {
		response := res.ErrorResponse(400, "unable to delete category", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// response
	response := res.SuccessResponse(200, "successfully category deleted", body)
	ctx.JSON(200, response)
}

// ----------------ViewCategory----------------

func (ct *ProductHandler) Viewcategory(ctx *gin.Context) {

	Paramid := ctx.Param("id")

	id, err := helper.StringToUInt(Paramid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find CategoryId",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	data, err := ct.ProductuseCase.FindCategoryById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't Find Category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "Category Found",
		Data:       data,
		Errors:     nil,
	})
}

// ----------------ViewFullCategory----------------

func (ct *ProductHandler) ViewFullcategory(ctx *gin.Context) {

	count, err1 := helper.StringToUInt(ctx.Query("count"))
	pageNumber, err2 := helper.StringToUInt(ctx.Query("page_number"))

	err1 = errors.Join(err1, err2)
	if err1 != nil {
		response := res.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var pagination req.ReqPagination
	pagination.Count = count
	pagination.PageNumber = pageNumber

	category, err := ct.ProductuseCase.ViewFullCategory(ctx, req.PageNation(pagination))

	if err != nil {
		response := res.ErrorResponse(500, "faild to get all category", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if category == nil {
		response := res.SuccessResponse(200, "there is no category to show", nil)
		ctx.JSON(http.StatusOK, response)
		return
	}

	respones := res.SuccessResponse(200, "successfully got all category", category)
	ctx.JSON(http.StatusOK, respones)

}
