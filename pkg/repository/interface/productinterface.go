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
	SaveCategory(ctx context.Context, category domain.Category) error
}
