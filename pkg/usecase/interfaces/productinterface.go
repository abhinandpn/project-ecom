package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type ProductuseCase interface {

	// Product
	AddProduct(ctx context.Context, product domain.Product) error                                // Add New Product
	UpdateProduct(ctx context.Context, product domain.Product) error                             // Update Product Info
	DeleteProduct(ctx context.Context, id uint) error                                            // Delete the product
	FindProductById(ctx context.Context, id uint) (res.ProductResponce, error)                   // Find a product by a ID
	ViewFullProduct(ctx context.Context, pagination req.PageNation) (res.ProductResponce, error) // View full Products

	// Category
	AddCategory(ctx context.Context, category req.CategoryReq) (cat res.CategoryRes, err error)
	Editcategory(ctx context.Context, category domain.Category) error
	DeleteCategory(ctx context.Context, categoryId uint) error
	ViewFullCategory(ctx context.Context, pagination req.PageNation) (category []res.CategoryRes, err error)
}
