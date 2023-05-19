package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type ProductRepository interface {

	// Product
	FindProductById(ctx context.Context, PId uint) (domain.Product, error)                         // Find The product Details By an ID
	ViewFullProduct(ctx context.Context, pagination req.PageNation) ([]res.ProductResponce, error) // View Full Product From database
	CreateProduct(ctx context.Context, product domain.Product) error                               // Save a New product in to database
	DeletProduct(ctx context.Context, PId uint) error                                              // Dalete product from database
	UpdateProduct(ctx context.Context, info domain.Product) (domain.Product, error)                // Update the product info
	// FindProduct(ctx context.Context, product domain.Product) (domain.Product, error)

	// Category managment
	FindCategoryById(ctx context.Context, CategoryId uint) (Category domain.Category, err error)
	FindAllCategory(ctx context.Context, pagination req.PageNation) ([]res.CategoryRes, error)
	SaveCategory(ctx context.Context, category req.CategoryReq) error
	UpdateCatrgoryName(ctx context.Context, category domain.Category) error
	DeletCategory(ctx context.Context, category domain.Category) error
}
