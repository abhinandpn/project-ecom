package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type ProductuseCase interface {

	// Product
	FindProductByName(ctx context.Context, name string) (domain.Product, error)                    // Find product by name
	FindProductById(ctx context.Context, id uint) (res.ProductResponce, error)                     // Find a product by a ID
	ViewFullProduct(ctx context.Context, pagination req.PageNation) ([]res.ProductResponce, error) // View full Products
	// curd
	AddProduct(ctx context.Context, product req.ReqProduct) error             // Add New Product
	UpdateProduct(ctx context.Context, product req.ReqProduct, id uint) error // Update Product Info
	DeleteProduct(ctx context.Context, id uint) error                         // Delete the product

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
	AddSubcategory(ctx context.Context, cid uint, name string) (domain.SubCategory, error)
	SubNameUpdate(ctx context.Context, id uint, name string) (domain.SubCategory, error)
	SubCatUpdate(ctx context.Context, id uint, name string) (domain.SubCategory, error)
	DeleteSubCate(ctx context.Context, name string) error
	ListALlSubCate(ctx context.Context, pagination req.PageNation) ([]res.SubCategoryRes, error)
	// finding
	FindSubById(ctx context.Context, id uint) (domain.SubCategory, error)
	FindSubByName(ctx context.Context, name string) (domain.SubCategory, error)
	FindSubByCatId(ctx context.Context, id uint) (domain.SubCategory, error)
	FindSubByCatName(ctx context.Context, name string) (domain.SubCategory, error)
}
