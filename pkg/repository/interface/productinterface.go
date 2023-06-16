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
}
