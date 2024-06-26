package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type ProductuseCase interface {

	// Category
	FindCategoryById(ctx context.Context, id uint) (res.CategoryRes, error)                     // Find category By an ID
	FindCategoryByname(ctx context.Context, name string) (domain.Category, error)               // Find category by name
	ViewFullCategory(ctx context.Context, pagination req.PageNation) ([]res.CategoryRes, error) // view full category
	// curd
	AddCategory(ctx context.Context, name string) (domain.Category, error)    // Add Cateory
	UpdateCategory(ctx context.Context, category req.UpdateCategoryReq) error // update category
	DeleteCategory(ctx context.Context, name string) error                    // Delete category

	// ------------------------SUB CATEGORY------------------------
	// curd
	AddSubcategory(ctx context.Context, cid uint, name string) (domain.SubCategory, error)       // add sub category
	SubNameUpdate(ctx context.Context, id uint, name string) (domain.SubCategory, error)         //sub category name update
	SubCatUpdate(ctx context.Context, id uint, name string) (domain.SubCategory, error)          // sub category update
	DeleteSubCate(ctx context.Context, name string) error                                        // delete sub category
	ListALlSubCate(ctx context.Context, pagination req.PageNation) ([]res.SubCategoryRes, error) // list all sub category
	// finding
	FindSubById(ctx context.Context, id uint) (domain.SubCategory, error)          // find sub categiry by id
	FindSubByName(ctx context.Context, name string) (domain.SubCategory, error)    // find sub categiry sub category  by name
	FindSubByCatId(ctx context.Context, id uint) (domain.SubCategory, error)       // find sub category by category id
	FindSubByCatName(ctx context.Context, name string) (domain.SubCategory, error) // find sb category by category name

	// ---------- PRODUCT UPDATED --------------

	// Finding
	GetProductByName(name string) (domain.Product, error)
	GetProductById(id uint) (domain.Product, error)
	GetAllProducts(pagination req.PageNation) ([]res.ProductResponce, error)
	GetAllQtyInfoProduct(pagination req.PageNation) ([]res.ProductQtyRes, error)
	ListproductByBrand(brand string) (domain.Product, error)
	ListProductByCategory(category string) (domain.Product, error)

	// product images
	AddImage(id uint, name string) error
	// ListProductBySubCategory(pagination req.PageNation, name string) ([]res.ProductResponce, error)

	// Opration
	CreateProduct(product req.ReqProduct) error
	DeleteProduct(id uint) error
	ProductUpdationNew(product req.UpdateProduct, id uint) error
	UpdateQuentity(id, qty uint) error

	// brand
	CreateBrand(name, img string) error
	DeleteBrand(id uint) error
	ViewFullBrand() ([]res.ResBrand, error)

	GetByColour(colour string, pagination req.PageNation) ([]res.ProductResponce, error)
	GetBySize(size int, pagination req.PageNation) ([]res.ProductResponce, error)
	GetByCategory(name string, pagination req.PageNation) ([]res.ProductResponce, error)
	GetByBrand(name string, pagination req.PageNation) ([]res.ProductResponce, error)
	GetByName(name string, pagination req.PageNation) ([]res.ProductResponce, error)
	GetByPrice(Start, End int, pagination req.PageNation) ([]res.ProductResponce, error)
	GetByQuantity(Start, End int, pagination req.PageNation) ([]res.ProductResponce, error)

	// final sort
	GetProductByString(name string, pagination req.PageNation) (res.ProductStringResponce, error)
}
