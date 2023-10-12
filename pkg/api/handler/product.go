package handler

import (
	"errors"
	"net/http"

	handlerInterface "github.com/abhinandpn/project-ecom/pkg/api/handler/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/helper"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductuseCase service.ProductuseCase
}

func NewProductHandler(productUsecase service.ProductuseCase) handlerInterface.ProductHandler {
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

	products, err := pr.ProductuseCase.GetAllQtyInfoProduct(req.PageNation(pagination))

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

	product := domain.Product{
		ProductName: body.ProductName,
		Discription: body.Discription,
		CategoryId:  body.CategoryID,
		BrandId:     body.BrandId,
		Info: domain.ProductInfo{
			Price:  float64(body.Price),
			Colour: body.Color,
			Size:   body.Size,
			// ImageId: body.ImageId,
		},
	}

	err := pr.ProductuseCase.CreateProduct(body)

	if err != nil {
		response := res.ErrorResponse(400, "faild to add product", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := res.SuccessResponse(200, "successfully product added", product)
	ctx.JSON(http.StatusOK, response)

}

func (pr *ProductHandler) UpdateProduct(ctx *gin.Context) {

	var body req.UpdateProduct
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

	err = pr.ProductuseCase.ProductUpdationNew(body, id)
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

	err = pr.ProductuseCase.DeleteProduct(id)
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
	data, err := pr.ProductuseCase.GetProductById(id)
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

// ------------------------sub category------------------------
// add sub category
func (sub *ProductHandler) AddSubCategory(ctx *gin.Context) {

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

	var body req.SubCateCurdRes
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	subcategory, err := sub.ProductuseCase.AddSubcategory(ctx, id, body.SubCategoryName)
	if err != nil {
		response := res.SuccessResponse(500, "failed to add subcategory", nil)
		ctx.JSON(http.StatusOK, response)
		return
	}
	respones := res.SuccessResponse(200, "successfully add sub category", subcategory)
	ctx.JSON(http.StatusOK, respones)

}

// delete sub category
func (sub *ProductHandler) DeleteSubCategory(ctx *gin.Context) {

	var body req.SubCateCurdRes
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	err := sub.ProductuseCase.DeleteSubCate(ctx, body.SubCategoryName)
	if err != nil {
		response := res.SuccessResponse(500, "failed to delete subcategory", nil)
		ctx.JSON(http.StatusOK, response)
		return
	}
	respones := res.SuccessResponse(200, "successfully delete sub category", nil)
	ctx.JSON(http.StatusOK, respones)

}

// view full sub category
func (sub *ProductHandler) ViewFullSubCategory(ctx *gin.Context) {

	count, err1 := helper.StringToUInt(ctx.Query("count"))
	pageNumber, err2 := helper.StringToUInt(ctx.Query("page_number"))

	err1 = errors.Join(err1, err2)
	if err1 != nil {
		response := res.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	var pagination req.PageNation
	pagination.Count = count
	pagination.PageNumber = pageNumber

	subcategory, err := sub.ProductuseCase.ListALlSubCate(ctx, pagination)

	if err != nil {
		response := res.ErrorResponse(500, "faild to get all sub category", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if subcategory == nil {
		response := res.SuccessResponse(200, "there is no sub category to show", nil)
		ctx.JSON(http.StatusOK, response)
		return
	}
	respones := res.SuccessResponse(200, "successfully get sub category", subcategory)
	ctx.JSON(http.StatusOK, respones)
}

// func (sub *ProductHandler) ViewSubCategory(ctx *gin.Context) {

// }

// update sub category
func (sub *ProductHandler) EditSubCategory(ctx *gin.Context) {

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
	var body req.SubCateCurdRes
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	subcate, err := sub.ProductuseCase.SubCatUpdate(ctx, id, body.SubCategoryName)
	if err != nil {
		response := res.ErrorResponse(500, "faild to update sub category", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	respones := res.SuccessResponse(200, "successfully get sub category", subcate)
	ctx.JSON(http.StatusOK, respones)
}

func (b *ProductHandler) AddBrand(ctx *gin.Context) {

	var body req.BrandReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := res.ErrorResponse(400, "invalid input", err.Error(), body)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	err := b.ProductuseCase.CreateBrand(body.BrandName, body.BrandImage)
	if err != nil {
		response := res.SuccessResponse(500, "failed to create brand", nil)
		ctx.JSON(http.StatusOK, response)
		return
	}
	respones := res.SuccessResponse(200, "successfully create brand", body)
	ctx.JSON(http.StatusOK, respones)
}

func (b *ProductHandler) DeletBrand(ctx *gin.Context) {

	Paramid := ctx.Param("id")

	id, err := helper.StringToUInt(Paramid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't brand id",
			Errors:     err.Error(),
			Data:       nil,
		})
		return
	}
	err = b.ProductuseCase.DeleteBrand(id)
	if err != nil {
		response := res.SuccessResponse(500, "failed to delete brand", nil)
		ctx.JSON(http.StatusOK, response)
		return
	}
	respones := res.SuccessResponse(200, "successfully delete brand", nil)
	ctx.JSON(http.StatusOK, respones)
}

func (b *ProductHandler) ViewBrands(ctx *gin.Context) {

	body, err := b.ProductuseCase.ViewFullBrand()
	if err != nil {
		response := res.SuccessResponse(500, "failed to get brands", nil)
		ctx.JSON(http.StatusOK, response)
		return
	}
	respones := res.SuccessResponse(200, "successfully delete brand", body)
	ctx.JSON(http.StatusOK, respones)
}

// ----------- Sorting -----------

func (p *ProductHandler) ProductGetByColour(ctx *gin.Context) {

	colour := ctx.Query("colour")
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

	products, err := p.ProductuseCase.GetByColour(colour, req.PageNation(pagination))

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

func (p *ProductHandler) ProductGetBySize(ctx *gin.Context) {

	count, err1 := helper.StringToUInt(ctx.Query("count"))
	pageNumber, err2 := helper.StringToUInt(ctx.Query("page_number"))
	size := ctx.Query("size")
	value, err3 := helper.StringToUInt(size)

	err1 = errors.Join(err1, err2, err3)
	if err1 != nil {
		response := res.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	pagination := req.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}

	products, err := p.ProductuseCase.GetBySize(int(value), req.PageNation(pagination))

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

func (p *ProductHandler) ProductGetByCategory(ctx *gin.Context) {

	count, err1 := helper.StringToUInt(ctx.Query("count"))
	pageNumber, err2 := helper.StringToUInt(ctx.Query("page_number"))
	category := ctx.Query("category")

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

	products, err := p.ProductuseCase.GetByCategory(category, req.PageNation(pagination))

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

func (p *ProductHandler) ProductGetByBrand(ctx *gin.Context) {

	count, err1 := helper.StringToUInt(ctx.Query("count"))
	pageNumber, err2 := helper.StringToUInt(ctx.Query("page_number"))
	brand := ctx.Query("brand")

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

	products, err := p.ProductuseCase.GetByBrand(brand, req.PageNation(pagination))

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

func (p *ProductHandler) ProductGetByName(ctx *gin.Context) {

	count, err1 := helper.StringToUInt(ctx.Query("count"))
	pageNumber, err2 := helper.StringToUInt(ctx.Query("page_number"))
	name := ctx.Query("name")

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

	products, err := p.ProductuseCase.GetByName(name, req.PageNation(pagination))

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

func (p *ProductHandler) ProductGetByPrice(ctx *gin.Context) {

	count, err1 := helper.StringToUInt(ctx.Query("count"))
	pageNumber, err2 := helper.StringToUInt(ctx.Query("page_number"))
	start := ctx.Query("start")
	end := ctx.Query("end")
	Pstart, err3 := helper.StringToUInt(start)
	Pend, err4 := helper.StringToUInt(end)

	err1 = errors.Join(err1, err2, err3, err4)
	if err1 != nil {
		response := res.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	pagination := req.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}

	products, err := p.ProductuseCase.GetByPrice(int(Pstart), int(Pend), req.PageNation(pagination))

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

func (p *ProductHandler) ProductGetByQuantity(ctx *gin.Context) {

	count, err1 := helper.StringToUInt(ctx.Query("count"))
	pageNumber, err2 := helper.StringToUInt(ctx.Query("page_number"))
	start := ctx.Query("start")
	end := ctx.Query("end")
	Qstart, err3 := helper.StringToUInt(start)
	Qend, err4 := helper.StringToUInt(end)

	err1 = errors.Join(err1, err2, err3, err4)
	if err1 != nil {
		response := res.ErrorResponse(400, "invalid inputs", err1.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	pagination := req.ReqPagination{
		PageNumber: pageNumber,
		Count:      count,
	}

	products, err := p.ProductuseCase.GetByQuantity(int(Qstart), int(Qend), req.PageNation(pagination))

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

func (p *ProductHandler) GetProductByString(ctx *gin.Context) {

	name := ctx.Query("name")
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

	data, err := p.ProductuseCase.GetProductByString(name, req.PageNation(pagination))
	if err != nil {
		response := res.ErrorResponse(500, "faild to get all products", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	respones := res.SuccessResponse(200, "successfully got all products", data)
	ctx.JSON(http.StatusOK, respones)
}
