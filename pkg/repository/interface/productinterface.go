package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type ProductRepository interface {
	// Product
	FindProduct(ctx context.Context, product domain.Product) (domain.Product, error)
	FindProductById(ctx context.Context, productId uint) (product domain.Product, err error)

	// List all product
	FindAllProduct(ctx context.Context, pagination req.PageNation) (products []res.ProductResponce, err error)

	// product Managment
	SaveProduct(ctx context.Context, product domain.Product) error
	UpdateProduct(ctx context.Context, product domain.Product) error

	// Category managment
	FindCategoryById(ctx context.Context, CategoryId uint) (Category domain.Category, err error)
	FindAllCategory(ctx context.Context, pagination req.PageNation) (category []res.CategoryRes, err error)
	SaveCategory(ctx context.Context, category domain.Category) error
	UpdateCatrgoryName(ctx context.Context, category domain.Category) error
	DeletCategory(ctx context.Context, category domain.Category) error
}
