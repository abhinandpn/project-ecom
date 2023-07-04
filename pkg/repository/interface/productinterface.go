package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type ProductRepository interface {

	// ------------------------Category------------------------
	FindCategoryById(ctx context.Context, CId uint) (domain.Category, error)                   // Find Category By Id
	FindCategoryByname(ctx context.Context, name string) (domain.Category, error)              // Find cTegory by name
	FindCategoryByName(name string) (domain.Category, error)                                   // Find Category By name
	FindAllCategory(ctx context.Context, pagination req.PageNation) ([]res.CategoryRes, error) // View Full category
	//------------CURD
	CreateCategory(ctx context.Context, name string) (domain.Category, error)                // Creating new Category
	UpdateCategory(ctx context.Context, body req.UpdateCategoryReq) (domain.Category, error) // Update category
	DeleteCategory(ctx context.Context, name string) (domain.Category, error)                // Delete category

	// ------------------------Subcategory------------------------
	// finding
	FindSubCtByName(ctx context.Context, name string) (domain.SubCategory, error)                   // Find by name
	FindSubCtById(ctx context.Context, id uint) (domain.SubCategory, error)                         // Find By id
	FindSubCtByCategoryName(ctx context.Context, name string) (domain.SubCategory, error)           // Find By Sub category name
	FindSubCtByCategoryId(ctx context.Context, id uint) (domain.SubCategory, error)                 // find by sub category id
	ListllSubCategory(ctx context.Context, pagination req.PageNation) ([]res.SubCategoryRes, error) // find all sub category
	// curd
	CreateSubCategory(ctx context.Context, cid uint, name string) (domain.SubCategory, error)       // create sub category
	DeleteSubCategory(ctx context.Context, name string) error                                       // delete sub category
	ChangeSubCatName(ctx context.Context, id uint, name string) (domain.SubCategory, error)         // change sub category name
	ChangeSubCatCatogeryName(ctx context.Context, id uint, name string) (domain.SubCategory, error) // change category for sub category

	// updated product query

	// findind
	FindProductByName(name string) (domain.Product, error)                             // Find Product By name
	FindProductById(id uint) (domain.Product, error)                                   // Find Product By Id
	FindProductByBrand(id uint) (domain.Product, error)                                // Find Product By Brand
	FindProductByCategory(id uint) (domain.Product, error)                             // Find Product By category
	FindProductBySubCat(id uint) (domain.Product, error)                               // Find Product By subcategory
	FindProductByProductInfo(id uint) (domain.Product, error)                          // Find Product By productinfo Id
	FindProductInfoById(id uint) (domain.ProductInfo, error)                           // Find Product info by Id
	FindAllProduct(pagination req.PageNation) ([]res.ProductResponce, error)           // Find All products
	ProductViewByPid(id uint) (res.ResProductOrder, error)                             // Find Produdct By Pinfo id
	FindAllProductWithQuantity(pagination req.PageNation) ([]res.ProductQtyRes, error) // Find All product with quantity

	// product images
	AddProductImage(id uint, img string) error
	FindImage(img string) (domain.ProductImage, error)

	// opration
	CreateProduct(product req.ReqProduct) error
	ProductUpdate(product req.UpdateProduct, id uint) error // Upatede product
	DeletProduct(id uint) error
	UpdateQuentity(id, qty uint) error

	// brand
	CreateBrand(name, img string) error
	DeleteBrand(id uint) error
	ViewFullBrand() ([]res.ResBrand, error)
	FindBrandByName(name string) (domain.Brand, error)
	FinBrandByName(name string) (domain.Brand, error) // Find Product By
	FIndBrandById(id uint) (domain.Brand, error)

	// ----------- Sorting -----------

	ListByColour(colour string, pagination req.PageNation) ([]res.ProductResponce, error)
	ListBySize(size uint, pagination req.PageNation) ([]res.ProductResponce, error)
	ListByCategory(id uint, pagination req.PageNation) ([]res.ProductResponce, error)
	ListByBrand(id uint, pagination req.PageNation) ([]res.ProductResponce, error)
	ListByName(name string, pagination req.PageNation) ([]res.ProductResponce, error)
	ListByPrice(Start, End float64, pagination req.PageNation) ([]res.ProductResponce, error)
	ListByQuantity(Start, End uint, pagination req.PageNation) ([]res.ProductResponce, error)
}
