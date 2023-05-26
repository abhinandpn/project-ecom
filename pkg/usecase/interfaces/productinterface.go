package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type ProductuseCase interface {

	// Product
	AddProduct(ctx context.Context, product domain.Product) error                                  // Add New Product
	UpdateProduct(ctx context.Context, product domain.Product) error                               // Update Product Info
	DeleteProduct(ctx context.Context, id uint) error                                              // Delete the product
	FindProductById(ctx context.Context, id uint) (res.ProductResponce, error)                     // Find a product by a ID
	ViewFullProduct(ctx context.Context, pagination req.PageNation) ([]res.ProductResponce, error) // View full Products

	// Category
	FindCategoryById(ctx context.Context, id uint) (res.CategoryRes, error)       // Find category By an ID
	FindCategoryByname(ctx context.Context, name string) (domain.Category, error) // Find category by name
	// CURD
	AddCategory(ctx context.Context, name string) (domain.Category, error)                      // Add Cateory
	UpdateCategory(ctx context.Context, category req.UpdateCategoryReq) error                   // update category
	ViewFullCategory(ctx context.Context, pagination req.PageNation) ([]res.CategoryRes, error) // view full category
	DeleteCategory(ctx context.Context, name string) error                                      // Delete category

}
