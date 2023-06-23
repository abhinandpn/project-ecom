package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type ProductRepository interface {

	/*
		------------------------Product------------------------
		FindProductById(ctx context.Context, PId uint) (domain.Product, error)                         // Find The product Details By an ID
		FindProductInfoByPid(ctx context.Context, pid uint) (domain.ProductInfo, error)                // findfind product info
		FindProductByName(ctx context.Context, name string) (domain.Product, error)                    // find product by name
		ViewFullProduct(ctx context.Context, pagination req.PageNation) ([]res.ProductResponce, error) // View Full Product From database
		UpdateQtyPinfo(ctx context.Context, pid uint, qty uint) error
		FindProductByPrinfo(pfid uint) (uint, error) // find product by product info
		------------curd
		CreateProduct(ctx context.Context, product req.ReqProduct) error       // Save a New product in to database
		DeletProduct(ctx context.Context, PId uint) error                      // Dalete product from database
		UpdateProduct(ctx context.Context, info req.ReqProduct, id uint) error // Update the product info
	*/
	// ------------------------Category------------------------
	FindCategoryById(ctx context.Context, CId uint) (domain.Category, error)                   // Find Category By Id
	FindCategoryByname(ctx context.Context, name string) (domain.Category, error)              // Find cTegory by name
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
	FindProductByName(name string) (domain.Product, error)
	FindProductById(id uint) (domain.Product, error)
	FindProductByBrand(id uint) (domain.Product, error)
	FindProductByCategory(id uint) (domain.Product, error)
	FindProductBySubCat(id uint) (domain.Product, error)
	FindAllProduct(pagination req.PageNation) ([]res.ProductResponce, error)
	FindProductByProductInfo(id uint) (domain.Product, error)
	FindProductWithQuentity(pagination req.PageNation) ([]res.ResProduct, error)
	FinBrandByName(name string) (domain.Brand, error)
	FindCategoryByName(name string) (domain.Category, error)
	// product images
	AddProductImage(id uint, img string) error
	FindImage(img string) (domain.ProductImage, error)

	// opration
	CreateProduct(product req.ReqProduct) error
	UpdateProduct(product res.ResProduct, id uint) error
	DeletProduct(id uint) error
	UpdateQuentity(id, qty uint) error

	// brand
	CreateBrand(name, img string) error
	DeleteBrand(id uint) error
	ViewFullBrand() (res.ResBrand, error)
	FindBrandByName(name string) (domain.Brand, error)
	FIndBrandById(id uint) (domain.Brand, error)
}
