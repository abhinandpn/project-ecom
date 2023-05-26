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

	// Category managment
	FindCategoryById(ctx context.Context, CId uint) (domain.Category, error)      // Find Category By Id
	FindCategoryByname(ctx context.Context, name string) (domain.Category, error) // Find cTegory by name
	// CURD
	CreateCategory(ctx context.Context, name string) (domain.Category, error)                  // Creating new Category
	DeleteCategory(ctx context.Context, id uint) error                                         // Delete category
	UpdateCategory(ctx context.Context, info domain.Category) (domain.Category, error)         // Update category
	FindAllCategory(ctx context.Context, pagination req.PageNation) ([]res.CategoryRes, error) // View Full category

}
