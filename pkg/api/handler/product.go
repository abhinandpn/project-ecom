package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/helper"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type ProductHandler struct {
	ProductuseCase service.ProductuseCase
}

func NewProductHandler(productUsecase service.ProductuseCase) *ProductHandler {
	return &ProductHandler{ProductuseCase: productUsecase}
}

/*
Products
 > Add
 > Edit
 > Delete
 > view (list product Admin and User Same)

*/

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
		fmt.Println("------------------------1 ")
		respones := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	product := domain.Product{
		ProductName: body.ProductName,
		CategoryID:  body.CategoryID,
		Discription: body.Discription,
		Price:       body.Price,
		Info: domain.ProductInfo{
			Colour: body.Color,
			Brand:  body.Brand,
			Size:   body.Size,
		},
	}

	fmt.Println("productdetails : ", body)

	err := pr.ProductuseCase.AddProduct(ctx, product)

	if err != nil {
		fmt.Println("------------------------2 ")

		response := res.ErrorResponse(400, "faild to add product", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := res.SuccessResponse(200, "successfully product added", product)
	ctx.JSON(http.StatusOK, response)

}

// todo category
// add
// edit
// delete
// update

func (pr *ProductHandler) Addcategory(ctx *gin.Context) {

	var body req.CategoryReq

	if err := ctx.ShouldBindJSON(&body); err != nil {
		fmt.Println("------------------------1 ")
		respones := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, respones)
		return
	}

	err := pr.ProductuseCase.AddCategory(ctx, body)
	if err != nil {
		response := res.ErrorResponse(400, "faild to add cateogy", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := res.SuccessResponse(200, "successfully added a new category", body)
	ctx.JSON(http.StatusOK, response)

}

func (pr *ProductHandler) EditCategory(ctx *gin.Context) {

	var body req.CategoryReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var category domain.Category
	copier.Copy(&category, &body)

	err := pr.ProductuseCase.Editcategory(ctx, category)
	if err != nil {
		response := res.ErrorResponse(400, "faild to update product", err.Error(), body)
		ctx.JSON(400, response)
		return
	}

	response := res.SuccessResponse(200, "successfully product updated", body)
	ctx.JSON(200, response)
}

func (pr *ProductHandler) DeleteCategory(ctx *gin.Context) {

	var body req.CategoryReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var category domain.Category
	copier.Copy(&category, &body)

	err := pr.ProductuseCase.DeleteCategory(ctx, category.CategoryID)
	if err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := res.SuccessResponse(200, "successfully product updated", body)
	ctx.JSON(200, response)

}

// todo Subcategory -- its updating
